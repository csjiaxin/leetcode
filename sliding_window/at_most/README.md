# atMost 相减：恰好 K 的子数组

[返回总目录](../../algorithm_templates/README.md)

[扩展练习：迁移点与边界](#extended-practice)

## 目录

- [1. 核心思路](#core)
- [2. 992：恰好 K 种整数](#lc992)
- [3. 1248：恰好 K 个奇数](#lc1248)
- [4. 930：二元数组和为 goal](#lc930)
- [5. 对比与易错点](#pitfalls)

<a id="core"></a>

## 1. 核心思路

```text
恰好 K 的数量 = 至多 K 的数量 - 至多 K-1 的数量

遍历右端点：
    加入当前元素
    while 超出预算：移除左端元素
    累加 right-left+1
```

**为什么加 right-left+1？** 收缩到最早合法 left 后，以 right 结尾、起点在 [left,right] 的所有子数组都合法，且更早的起点都非法。

公式来自集合相减；滑窗实现另有条件：删掉左端元素不能让约束变坏。种数、奇数数目、非负数组的和满足要求，任意负数数组的和不满足。

**通用的是循环步骤，不需要 add/remove/valid 回调。** 每次 atMost 调用都建立独立窗口；k<0 返回 0。下面每题都有独立、完整的实现，不依赖其他章节代码。代码仅在文档中，练习文件保持原样。

<a id="lc992"></a>

## 2. [992. K 个不同整数的子数组](https://leetcode.com/problems/subarrays-with-k-different-integers/)

[练习](l992/992.go) · [测试](l992/992_test.go)

**题意**：统计恰好包含 k 种不同整数的非空连续子数组。

**状态与不变量**：freq 保存窗口内频次，distinct 保存不同数的数量。收缩后 distinct<=limit；频次从0变1时增加种数，从1变0时减少种数。

```go
package template

func subarraysWithKDistinct(nums []int, k int) int {
	atMost := func(limit int) int64 {
		if limit < 0 {
			return 0
		}
		freq := make(map[int]int)
		left, distinct := 0, 0
		var answer int64
		for right, x := range nums {
			if freq[x] == 0 {
				distinct++
			}
			freq[x]++
			for distinct > limit {
				y := nums[left]
				freq[y]--
				if freq[y] == 0 {
					delete(freq, y)
					distinct--
				}
				left++
			}
			answer += int64(right - left + 1)
		}
		return answer
	}
	return int(atMost(k) - atMost(k-1))
}
```

**推演**：[1,2,1]、k=2。atMost(2)=1+2+3=6；atMost(1)=3，故答案为3，分别是 [1,2]、[2,1]、[1,2,1]。

**复杂度**：时间 O(n)，空间 O(min(n,k+1))。两次扫描仍为线性；删除零频键避免 map 保留无用历史值。

<a id="lc1248"></a>

## 3. [1248. 统计优美子数组](https://leetcode.com/problems/count-number-of-nice-subarrays/)

[练习](l1248/1248.go) · [测试](l1248/1248_test.go)

**题意**：统计恰好包含 k 个奇数的非空连续子数组。

**状态**：只需 odd，不再需要频次表。奇数贡献1，偶数贡献0；加入和移除时维护 odd。

```go
package template

func numberOfSubarrays(nums []int, k int) int {
	atMost := func(limit int) int64 {
		if limit < 0 {
			return 0
		}
		left, odd := 0, 0
		var answer int64
		for right, x := range nums {
			if x%2 != 0 {
				odd++
			}
			for odd > limit {
				if nums[left]%2 != 0 {
					odd--
				}
				left++
			}
			answer += int64(right - left + 1)
		}
		return answer
	}
	return int(atMost(k) - atMost(k-1))
}
```

**例子**：[2,1,2]、k=1，答案4。唯一奇数的左右两边都可选择是否带上偶数，不能只对每个右端点加1。

**复杂度**：时间 O(n)，空间 O(1)。另一种解法是 [前缀奇数数量＋哈希表](../../Subarray/readme.md#lc1248)。

<a id="lc930"></a>

## 4. [930. 和相同的二元子数组](https://leetcode.com/problems/binary-subarrays-with-sum/)

[练习](l930/930.go) · [测试](l930/930_test.go)

**题意**：数组仅含0、1，统计和恰好为 goal 的非空子数组。

**状态**：sum 是窗口和。非负元素确保删左端不会增加和，因此可以用至多和的窗口。

```go
package template

func numSubarraysWithSum(nums []int, goal int) int {
	atMost := func(limit int) int64 {
		if limit < 0 {
			return 0
		}
		left, sum := 0, 0
		var answer int64
		for right, x := range nums {
			sum += x
			for sum > limit {
				sum -= nums[left]
				left++
			}
			answer += int64(right - left + 1)
		}
		return answer
	}
	return int(atMost(goal) - atMost(goal-1))
}
```

**例子**：[0,0,0]、goal=0。atMost(0)=1+2+3=6；atMost(-1)=0，答案6。这就是负预算提前返回的重要性。

**复杂度**：时间 O(n)，空间 O(1)。含负数的普通数组不可照搬，可改用前缀和计数。

<a id="pitfalls"></a>

## 5. 对比与易错点

| 题目 | 窗口状态 | 违规条件 | 每轮贡献 |
| --- | --- | --- | --- |
| 992 | freq、distinct | distinct>limit | right-left+1 |
| 1248 | odd | odd>limit | right-left+1 |
| 930 | sum | sum>limit | right-left+1 |

- 求的是数量，累加所有合法后缀，不是维护最长长度。
- 每次 atMost 都重置 left 和状态，不能沿用第一次的状态。
- 窗口清空时 left=right+1，贡献为0，不计空子数组。
- 计数中间值用 int64；上述题目答案在 LeetCode int 范围内，最后转 int。
- 所有独立代码块可单独复制，不能把多个 `package template` 合到同一文件。

<a id="extended-practice"></a>

## 扩展练习：迁移点与边界

新增题目只补充学习路线与解法总结，不自动生成未请求的实现或测试。

| 题目（原题） | 如何迁移当前模板 | 必须确认的条件 / 边界 |
| --- | --- | --- |
| [2062. 统计字符串中的元音子字符串](https://leetcode.com/problems/count-vowel-substrings-of-a-string/) | 仅元音的段内恰好五种=atMost(5)-atMost(4)，辅音出现时重置窗口 | 不能允许窗口跨过辅音 |
| [3306. 元音辅音字符串计数II](https://leetcode.com/problems/count-of-substrings-containing-every-vowel-and-k-consonants-ii/) | 全元音至少一次且恰好k辅音；可用atLeast(k)-atLeast(k+1) | 同时维护元音覆盖和辅音数，不能只数辅音 |
| [1358. 包含所有三种字符的子字符串数目](https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/) | 字符仅a/b/c，可用全部区间-atMost(2)，或覆盖后计数 | 这是至少覆盖而非直接恰好k事件的普通改名 |
| [2537. 统计好子数组的数目](https://leetcode.com/problems/count-the-number-of-good-subarrays/) | 加入值x新增freq[x]个相等对，维护至少k对窗口 | 对数不是不同值种数；移除顺序影响减少的对数 |
| [2302. 统计得分小于K的子数组数目](https://leetcode.com/problems/count-subarrays-with-score-less-than-k/) | 正数窗口得分=和×长度，收缩到<k后累加合法后缀 | 严格小于；乘法和答案用int64 |
| [2762. 不间断子数组](https://leetcode.com/problems/continuous-subarrays/) | 合法条件max-min<=2，单调双端队列维护最值，合法后累加长度 | 仍是至多型计数，但不能只用频次数量判断合法 |

**复习方法**：先独立写出状态、合法性或松弛条件，再指出相较本篇基础题改变的那一行；若前提不满足，主动换算法，不为统一模板牺牲正确性。
