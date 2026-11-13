# 双指针（Two Pointers）

双指针的核心不是“定义两个下标”，而是利用顺序、单调性或结构关系，让指针协同移动，避免重复枚举。

题单入口：[LeetCode Two Pointers](https://leetcode.com/problem-list/two-pointers/)。本文按经典模型整理代表题，不是该动态题单的完整枚举。

## 目录

- [1. 核心思想与模型速查](#core)
- [2. 相向双指针](#opposite)
- [3. 同向读写指针](#read-write)
- [4. 双序列归并与匹配](#merge)
- [5. 快慢指针与滑动窗口](#fast-slow)
- [6. 15. 三数之和：模板与图解](#leetcode-15)
- [7. 11. 盛最多水的容器：模板与图解](#leetcode-11)
- [8. 42. 接雨水：模板与图解](#leetcode-42)
- [9. 142. 环形链表 II：模板与图解](#leetcode-142)
- [10. 分区指针、易错点与练习路线](#practice)
- [11. 原参考链接（保留）](#references)

<a id="core"></a>

## 1. 核心思想与模型速查

做题前先回答三个问题：

1. 每个指针代表什么，维护的区间是什么？
2. 什么条件下移动哪个指针？
3. **为什么移动后不会漏解？**

| 模型 | 指针含义 | 常见信号 | 代表题 |
|---|---|---|---|
| 相向双指针 | 左右边界向中间移动 | 有序数对、回文、两端高度 | 167、15、11、42、125 |
| 同向读写指针 | 一个扫描，一个维护有效结果 | 原地删除、去重、移动元素 | 26、27、80、283 |
| 双序列指针 | 分别扫描两个序列 | 合并、比较、匹配 | 88、392、844、986 |
| 快慢 / 间距指针 | 不同速度或固定距离 | 中点、环、倒数位置 | 876、141、142、19 |
| 滑动窗口 | 连续区间的左右边界 | 子数组、子串、区间约束 | 3、209、713、76 |
| 分区与多指针 | 各类别的边界 | 原地分类、重排 | 75、905 |

分类可以重叠，重点是识别不变量。两个指针各自最多单向移动 n 次时，总移动次数是 O(n)，不是 O(n²)；排序、外层枚举等成本需要另外计算。

<a id="opposite"></a>

## 2. 相向双指针

### 2.1 有序数对：根据大小排除候选

| 题目 | 核心思路 | 时间（包含排序） |
|---|---|---|
| [167. 两数之和 II](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/) | 和小移左，和大移右 | O(n) |
| [15. 三数之和](https://leetcode.com/problems/3sum/) | 排序，固定一个数，其余做两数之和 | O(n²) |
| [16. 最接近的三数之和](https://leetcode.com/problems/3sum-closest/) | 固定一个数，维护最小差值 | O(n²) |
| [18. 四数之和](https://leetcode.com/problems/4sum/) | 固定两个数，其余相向查找 | O(n³) |
| [611. 有效三角形的个数](https://leetcode.com/problems/valid-triangle-number/) | 固定最大边，一次统计一段合法左端点 | O(n²) |

167 的排除依据（数组升序）：

- `nums[l]+nums[r] < target`：当前 l 与任何不超过 r 的右端点搭配都不够大，排除 l。
- `nums[l]+nums[r] > target`：当前 r 与任何不小于 l 的左端点搭配都太大，排除 r。

611：排序后固定最大边 `nums[k]`，令 `l=0, r=k-1`。若 `nums[l]+nums[r]>nums[k]`，则 l 到 r-1 都可作为第一条边，一次贡献 `r-l`，然后 `r--`；否则 `l++`。

### 2.2 回文、反转与两端选择

| 题目 | 核心思路 |
|---|---|
| [344. 反转字符串](https://leetcode.com/problems/reverse-string/) | 两端交换，向中间收缩 |
| [345. 反转字符串中的元音字母](https://leetcode.com/problems/reverse-vowels-of-a-string/) | 两端各自寻找元音，再交换 |
| [125. 验证回文串](https://leetcode.com/problems/valid-palindrome/) | 跳过非字母数字字符，忽略大小写后比较 |
| [680. 验证回文串 II](https://leetcode.com/problems/valid-palindrome-ii/) | 首次失配时检查跳过左字符或右字符 |
| [977. 有序数组的平方](https://leetcode.com/problems/squares-of-a-sorted-array/) | 比较两端绝对值，从结果末尾填入最大平方 |

680 不能随便删除一边：需要验证 `[l+1,r]` 或 `[l,r-1]` 是否为回文。最多额外做两次线性检查，总时间仍为 O(n)。

<a id="read-write"></a>

## 3. 同向读写指针

**读指针扫描输入，写指针维护有效前缀。** 不变量：`nums[:write]` 是已处理部分的正确结果。

| 题目 | 保留条件 / 操作 |
|---|---|
| [27. 移除元素](https://leetcode.com/problems/remove-element/) | 保留不等于 val 的元素 |
| [26. 删除有序数组中的重复项](https://leetcode.com/problems/remove-duplicates-from-sorted-array/) | 与最后保留的元素不同才写入 |
| [80. 删除有序数组中的重复项 II](https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/) | 每个值最多保留两次 |
| [283. 移动零](https://leetcode.com/problems/move-zeroes/) | 稳定写入非零元素，再把尾部补零 |
| [443. 压缩字符串](https://leetcode.com/problems/string-compression/) | 读取连续字符组，写入字符与计数 |

### 27：基本模板

```go
func removeElement(nums []int, val int) int {
	write := 0
	for _, x := range nums {
		if x != val {
			nums[write] = x
			write++
		}
	}
	return write
}
```

返回有效长度，尾部内容不重要。时间 O(n)，额外空间 O(1)。

### 26 / 80：统一为最多保留 k 次

```go
// 前提：nums 有序，k >= 1。结果保存在 nums[:write]。
func keepAtMostK(nums []int, k int) int {
	write := 0
	for _, x := range nums {
		if write < k || x != nums[write-k] {
			nums[write] = x
			write++
		}
	}
	return write
}
```

`k=1` 对应 26，`k=2` 对应 80。比较的是已保留结果中的 `nums[write-k]`，不是原数组中读指针之前第 k 个元素。写入位置不会超过当前读取位置，因此不会破坏尚未读取的数据。

<a id="merge"></a>

## 4. 双序列归并与匹配

| 题目 | 核心思路 | 时间 |
|---|---|---|
| [88. 合并两个有序数组](https://leetcode.com/problems/merge-sorted-array/) | 从后往前填，避免覆盖未读元素 | O(m+n) |
| [392. 判断子序列](https://leetcode.com/problems/is-subsequence/) | 长串持续扫描，匹配成功才推进短串 | O(m+n) |
| [844. 比较含退格的字符串](https://leetcode.com/problems/backspace-string-compare/) | 从后向前扫描，各自跳过被删除字符 | O(m+n) |
| [986. 区间列表的交集](https://leetcode.com/problems/interval-list-intersections/) | 计算交集，推进结束更早的区间 | O(m+n) |

986：当前两个闭区间的交集为 `[max(a.start,b.start), min(a.end,b.end)]`，左端点不大于右端点时存在交集。之后推进结束更早的区间，因为它不可能再与另一序列后面的区间相交。前提是每组区间各自有序且内部不相交。

<a id="fast-slow"></a>

## 5. 快慢指针与滑动窗口

### 5.1 快慢 / 固定间距指针

| 题目 | 指针设计 | 关键点 |
|---|---|---|
| [876. 链表的中间结点](https://leetcode.com/problems/middle-of-the-linked-list/) | slow 一步，fast 两步 | 常见模板在偶数长度时返回第二个中点 |
| [141. 环形链表](https://leetcode.com/problems/linked-list-cycle/) | slow 一步，fast 两步 | 环内一定相遇 |
| [142. 环形链表 II](https://leetcode.com/problems/linked-list-cycle-ii/) | 相遇后，一个回头部，双方每次一步 | 再次相遇即环入口 |
| [19. 删除链表的倒数第 N 个结点](https://leetcode.com/problems/remove-nth-node-from-end-of-list/) | 哑节点配合固定间距 | 最终让慢指针停在待删节点前驱 |
| [234. 回文链表](https://leetcode.com/problems/palindrome-linked-list/) | 找中点、反转后半段、比较 | 注意奇偶长度，需要时恢复链表 |
| [202. 快乐数](https://leetcode.com/problems/happy-number/) | 将数位平方和视为后继函数 | 判断到达 1 或进入其他环 |
| [287. 寻找重复数](https://leetcode.com/problems/find-the-duplicate-number/) | 下标到数值形成函数图 | 依赖题目数值范围，转为找环入口 |

快慢指针不限于链表：存在确定后继函数 `x → f(x)` 时，也可以研究轨迹是否入环。

### 5.2 滑动窗口

滑动窗口维护连续区间，是双指针的重要分支。关键前提是状态可增量维护，并且指针移动具有可利用的单调性。

| 模型 | 代表题 | 收缩 / 统计方式 |
|---|---|---|
| 最长合法区间 | [3](https://leetcode.com/problems/longest-substring-without-repeating-characters/)、[424](https://leetcode.com/problems/longest-repeating-character-replacement/)、[1004](https://leetcode.com/problems/max-consecutive-ones-iii/) | 不合法时收缩 |
| 最短满足区间 | [209](https://leetcode.com/problems/minimum-size-subarray-sum/)、[76](https://leetcode.com/problems/minimum-window-substring/) | 满足时更新答案并继续收缩 |
| 合法区间计数 | [713](https://leetcode.com/problems/subarray-product-less-than-k/) | 调整到合法后累计 `right-left+1` |
| 恰好 K 个不同值 | [992](https://leetcode.com/problems/subarrays-with-k-different-integers/) | `atMost(K)-atMost(K-1)` |
| 定长窗口 | [438](https://leetcode.com/problems/find-all-anagrams-in-a-string/)、[567](https://leetcode.com/problems/permutation-in-string/) | 保持长度并维护频次 |

209 的常规窗口依赖正数条件；允许负数后不能直接套用。713 的常规乘积窗口依赖正整数条件，并需处理 `k<=1`。详见[滑动窗口总结](../sliding_window/README.md)。

<a id="leetcode-15"></a>

## 6. [15. 三数之和](https://leetcode.com/problems/3sum/)：模板与图解

### 思路与移动依据

先排序，固定 `nums[i]`，在 `[i+1,n-1]` 中寻找和为 `-nums[i]` 的两个数。

- 三数之和小于 0：当前 l 即使配上当前最大的 r 也不够大，`l++`。
- 三数之和大于 0：当前 r 即使配上当前最小的 l 也太大，`r--`。
- 等于 0：记录答案，移动两端并跳过重复值。
- `nums[i]>0` 时后面全为正数，可以停止。

### Go 模板

```go
import "sort"

// 会原地排序 nums；需要保留输入时，先复制切片。
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i+2 < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			// 在加法前转换，避免先用 int 求和再转换。
			sum := int64(nums[i]) + int64(nums[l]) + int64(nums[r])
			switch {
			case sum < 0:
				l++
			case sum > 0:
				r--
			default:
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}
	}
	return ans
}
```

时间 O(n²)；双指针扫描自身额外空间 O(1)，排序栈空间随实现而定，返回结果另计。`int64` 足够覆盖本题数值约束，不代表任意三个 int64 都不会溢出。

### 逐步图解

```text
输入：[-1, 0, 1, 2, -1, -4]
排序：[-4,-1,-1, 0, 1, 2]
下标： 0  1  2  3  4  5
```

| 固定 i | l | r | 三数之和 | 操作 |
|---|---|---|---|---|
| 0 (-4) | 1 (-1) | 5 (2) | -3 | l++ |
| 0 (-4) | 2 (-1) | 5 (2) | -3 | l++ |
| 0 (-4) | 3 (0) | 5 (2) | -2 | l++ |
| 0 (-4) | 4 (1) | 5 (2) | -1 | l++，本轮结束 |
| 1 (-1) | 2 (-1) | 5 (2) | 0 | 记录 [-1,-1,2]，两端内移 |
| 1 (-1) | 3 (0) | 4 (1) | 0 | 记录 [-1,0,1]，两端内移 |
| 2 (-1) | — | — | — | 与前一个固定值相同，跳过 |
| 3 (0) | 4 (1) | 5 (2) | 3 | r--，本轮结束 |

去重示例：`[-2,0,0,2,2]` 固定 -2，第一次找到 `[-2,0,2]` 后，l/r 内移并跳过重复值，不能再输出同一个值三元组。

**边界：** 长度不足 3 返回空结果；`[0,0,0,0]` 只输出一个 `[0,0,0]`；三个元素必须来自不同下标。

<a id="leetcode-11"></a>

## 7. [11. 盛最多水的容器](https://leetcode.com/problems/container-with-most-water/)：模板与图解

### 思路与证明：为什么移动短板？

`area(l,r) = min(height[l],height[r]) × (r-l)`。

假设 `height[l] <= height[r]`，固定 l，把右端点换成任意 `r'<r`：

```text
area(l,r') <= height[l] × (r'-l)
           <= height[l] × (r-l)
           =  area(l,r)
```

因此当前 l 不可能再参与更优答案，可以 `l++`。右边更矮时对称。两端等高时移动任意一边即可。

**移动短板只是保留可能变好的机会，不保证下一步面积一定增大。**

### Go 模板

```go
func maxArea(height []int) int {
	l, r, best := 0, len(height)-1, 0
	for l < r {
		h := height[l]
		if height[r] < h {
			h = height[r]
		}
		area := h * (r - l)
		if area > best {
			best = area
		}
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return best
}
```

时间 O(n)，额外空间 O(1)。本题约束下 int 可容纳面积；推广到更大数值范围时应先转换再用 int64 乘法。

### 逐步图解

```text
下标：  0 1 2 3 4 5 6 7 8
高度： [1,8,6,2,5,4,8,3,7]
        L               R   → 面积 1×8=8，移动左侧短板
          L             R   → 面积 7×7=49，移动右侧短板
          L           R     → 面积 3×6=18，继续移动右侧
```

| l | r | 短板高度 | 宽度 | 面积 | 下一步 |
|---|---|---|---|---|---|
| 0 | 8 | 1 | 8 | 8 | l++ |
| 1 | 8 | 7 | 7 | 49 | r-- |
| 1 | 7 | 3 | 6 | 18 | r-- |
| 1 | 6 | 8 | 5 | 40 | 等高，模板选择 l++ |
| 2 | 6 | 6 | 4 | 24 | l++ |
| 3 | 6 | 2 | 3 | 6 | l++ |
| 4 | 6 | 5 | 2 | 10 | l++ |
| 5 | 6 | 4 | 1 | 4 | l++，结束 |

最大面积为 49。中间柱子不会减少容器容量；本题不是接雨水。

**边界：** 全零答案为 0；两个元素只计算一次；模板对空数组 / 单元素返回 0。

<a id="leetcode-42"></a>

## 8. [42. 接雨水](https://leetcode.com/problems/trapping-rain-water/)：模板与图解

### 从公式到双指针

位置 i 的水量为 `min(左侧最高柱,右侧最高柱)-height[i]`，这里两侧最高值都包含当前位置，水量不会为负数。

维护 `[0,l]` 的最高柱 leftMax 与 `[r,n-1]` 的最高柱 rightMax。每轮先用当前位置更新它们：

- `leftMax <= rightMax`：右边已经存在不低于 leftMax 的柱子，l 位置水位确定为 leftMax，结算后 `l++`。
- 否则：左边已经存在不低于 rightMax 的柱子，r 位置水位确定为 rightMax，结算后 `r--`。

中间未处理区域可能出现更高柱子，但不会改变较低侧当前被结算位置的水位。**不是猜另一侧最终最高值，而是已有足够高的挡板。**

### Go 模板

```go
func trap(height []int) int {
	l, r := 0, len(height)-1
	leftMax, rightMax, water := 0, 0, 0
	for l <= r {
		if height[l] > leftMax {
			leftMax = height[l]
		}
		if height[r] > rightMax {
			rightMax = height[r]
		}
		if leftMax <= rightMax {
			water += leftMax - height[l]
			l++
		} else {
			water += rightMax - height[r]
			r--
		}
	}
	return water
}
```

时间 O(n)，额外空间 O(1)。`l<=r` 表示每个位置结算一次；空数组不会进入循环。不要把其他变体的比较对象、最高值更新时机和循环条件随意拼接。

### 逐步图解

选择能同时展示左右结算的例子 `[4,2,0,3,2,5,1]`：

```text
#：柱子，~：积水

高度 5 | . . . . . # .
高度 4 | # ~ ~ ~ ~ # .
高度 3 | # ~ ~ # ~ # .
高度 2 | # # ~ # # # .
高度 1 | # # ~ # # # #
下标   | 0 1 2 3 4 5 6
水量   | 0 2 4 1 2 0 0  → 总水量 9
```

表中最高值均为本轮更新之后的值：

| l | r | leftMax | rightMax | 结算位置 | 新增水量 | 累计 |
|---|---|---|---|---|---|---|
| 0 | 6 | 4 | 1 | 右侧 6 | 1-1=0 | 0 |
| 0 | 5 | 4 | 5 | 左侧 0 | 4-4=0 | 0 |
| 1 | 5 | 4 | 5 | 左侧 1 | 4-2=2 | 2 |
| 2 | 5 | 4 | 5 | 左侧 2 | 4-0=4 | 6 |
| 3 | 5 | 4 | 5 | 左侧 3 | 4-3=1 | 7 |
| 4 | 5 | 4 | 5 | 左侧 4 | 4-2=2 | 9 |
| 5 | 5 | 5 | 5 | 左侧 5 | 5-5=0 | 9 |

第二轮开始，右侧已经发现高度 5 的柱子，因此左边高度上限为 4 的各位置都可以安全结算。

### 与 11 的区别

| 对比 | 11 盛最多水的容器 | 42 接雨水 |
|---|---|---|
| 目标 | 选一对边界使面积最大 | 累加每个位置的积水 |
| 状态 | 两端高度与宽度 | 两侧最高柱 |
| 移动依据 | 排除不可能更优的端点 | 结算水位已经确定的位置 |
| 中间柱子 | 不影响这对边界的容量计算 | 占据空间，必须扣除柱高 |

**边界：** 单调数组、全相等数组均不积水；少于三根柱子答案为 0；不要只看左右紧邻柱子，要看两侧最高柱。

<a id="leetcode-142"></a>

## 9. [142. 环形链表 II](https://leetcode.com/problems/linked-list-cycle-ii/)：模板与图解

### 两个阶段

1. 判环：slow 每次一步，fast 每次两步；fast 或 fast.Next 为空说明无环。
2. 找入口：相遇后，让一个指针回到 head，另一个留在相遇点；两者每次各走一步，再次相遇的位置就是入口。

### Go 模板

```go
// LeetCode 已提供 ListNode，提交时不要重复定义。
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			entry := head
			for entry != slow {
				entry = entry.Next
				slow = slow.Next
			}
			return entry
		}
	}
	return nil
}
```

时间 O(n)，额外空间 O(1)。判断节点是否相同必须比较指针，不能比较 Val；不同节点的值可能相同。题目中的 pos 只是描述尾节点连接位置，不是函数入参。

### 逐步图解

```text
A → B → C → D → E
        ↑       ↓
        └───────┘

入口 C；头部到入口距离 a=2；环长 c=3。
```

阶段一（初始均在 A，先移动再判断）：

| 轮次 | slow（一步） | fast（两步） | 结果 |
|---|---|---|---|
| 初始 | A | A | 初始相同不能直接判有环 |
| 1 | B | C | 未相遇 |
| 2 | C | E | 未相遇 |
| 3 | D | D | 相遇，但 D 不是入口 |

阶段二：entry 回到 A，slow 留在 D：

| 轮次 | entry | slow | 结果 |
|---|---|---|---|
| 初始 | A | D | 每次各走一步 |
| 1 | B | E | 未相遇 |
| 2 | C | C | 找到入口 C |

### 为什么回到头部后会在入口相遇？

设头部到入口距离为 a，入口沿环到相遇点距离为 b（`0<=b<c`），环长为 c。slow 到相遇时可能已经绕环 k 圈：

```text
slow 距离 s = a + b + k*c
fast 距离   = 2*s
距离差 s    = m*c       （m >= 1）
因此 a+b ≡ 0 (mod c)
     a   ≡ -b (mod c)
```

从相遇点沿环走 a 步，最终位置是 `(b+a) mod c = 0`，恰好回到入口；从 head 走 a 步也到入口。

若 a>0，head 出发的指针在前 a 步尚未进入环，而另一个指针一直在环内，因此不可能提前相遇；若 a=0，两者在第二阶段开始时就处于入口。

**边界：** 空链表、单节点无环返回 nil；单节点自环返回自身；入口可能为 head；不能把第一次相遇点直接当作入口。

<a id="practice"></a>

## 10. 分区指针、易错点与练习路线

### 10.1 75：荷兰国旗分区

[75. 颜色分类](https://leetcode.com/problems/sort-colors/) 用三个指针维护四个区间：

```text
[0,low)     全是 0
[low,i)     全是 1
[i,high]    尚未处理
(high,n)    全是 2
```

```go
func sortColors(nums []int) {
	low, i, high := 0, 0, len(nums)-1
	for i <= high {
		switch nums[i] {
		case 0:
			nums[low], nums[i] = nums[i], nums[low]
			low++
			i++
		case 1:
			i++
		case 2:
			nums[i], nums[high] = nums[high], nums[i]
			high--
		}
	}
}
```

时间 O(n)，额外空间 O(1)。前提是元素只有 0、1、2。**交换到右边之后不执行 i++**，因为换回来的元素尚未处理。

### 10.2 常见错误

1. **没有证明就移动指针**：移动应对应“排除不可能的候选”或“完成某个位置的处理”。
2. **忽略排序影响**：排序会改变原始下标，并产生时间与栈空间成本；需要原下标时应额外保存信息。
3. **混淆循环边界**：不同元素配对通常 `l<r`；逐点结算或处理未知区间时可能 `l<=r`，应由不变量决定。
4. **去重不完整**：15 / 18 的固定值与匹配成功后的左右值都要考虑去重。
5. **原地覆盖未读数据**：同向写入通常保证 write 不超过 read；88 则从后往前填。
6. **Go 字符串字节与字符混淆**：`s[i]` 返回字节；ASCII 题可直接处理，按 Unicode 码点处理时考虑 `[]rune`。
7. **看到连续区间就套窗口**：先检查约束是否支持单调收缩，尤其注意负数。
8. **整数溢出**：先转为 int64 再求和 / 乘法，不是算完后再转换；仍需结合数值范围确认安全。
9. **链表比较节点值**：环检测比较节点身份（指针），不是值。

### 10.3 推荐刷题顺序

| 阶段 | 题目 | 练习目标 |
|---|---|---|
| 基础移动 | 344 → 125 → 27 → 26 → 283 | 边界、扫描、有效前缀 |
| 相向排除 | 167 → 15 → 16 → 18 → 611 | 单调性、去重、批量计数 |
| 贪心证明 | 11 → 42 | 区分排除候选与确定水位 |
| 双序列协作 | 88 → 392 → 844 → 986 | 选择移动哪一侧 |
| 链表与函数图 | 876 → 141 → 142 → 19 → 234 → 287 | 速度差、固定间距、环入口 |
| 综合应用 | 80 → 75 → 3 → 209 → 713 → 76 → 992 | 不变量与窗口条件 |

**一句话记忆：相向指针排除候选，读写指针维护结果，归并指针推进序列，快慢指针利用距离，滑动窗口维护区间。**

<a id="references"></a>

## 11. 原参考链接（保留）

- 综合刷题经验：https://leetcode.com/discuss/study-guide/1688903/Solved-all-two-pointers-problems-in-100-days.

RainWater：

- https://leetcode.com/problems/trapping-rain-water/submissions/
- https://leetcode.com/problems/container-with-most-water/
- https://leetcode.com/problems/trapping-rain-water/discuss/17357/Sharing-my-simple-c%2B%2B-code%3A-O(n)-time-O(1)-space
- https://leetcode.com/problems/trapping-rain-water/discuss/153992/Java-O(n)-time-and-O(1)-space-(with-explanations).

原有实现：[11.go](rainwater/11.go)、[42.go](rainwater/42.go)。本文模板与原实现可能采用不同的等高处理或循环边界，不影响各自的不变量；本次只更新学习笔记，不修改原实现。
