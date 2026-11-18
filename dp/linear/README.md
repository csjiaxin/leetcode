# 线性 DP：前缀依赖、相邻选择与前缀最优

[返回 DP 总目录](../README.md)

同样按下标推进，状态可以是方案数、以当前位置结尾的最优值、或少量有限状态。先确定信息含义，再选以下三个子模板。

| 子模板 | 适用条件与不变量 | 题目 |
| --- | --- | --- |
| PrefixTable | dp[0] 为空状态；计算 i 时只读取下标小于 i 的状态 | [70](https://leetcode.com/problems/climbing-stairs/) 计数、[300](https://leetcode.com/problems/longest-increasing-subsequence/) 结尾长度、[343](https://leetcode.com/problems/integer-break/) 拆分乘积 |
| Independent | 更新右侧全部读取旧 skip/take；相邻项不能同时选 | [198](https://leetcode.com/problems/house-robber/)；[213](https://leetcode.com/problems/house-robber-ii/) 拆成两个线性区间 |
| PrefixBest | 当前答案先查询严格更早的候选，再把新候选纳入历史最优 | [1014](https://leetcode.com/problems/best-sightseeing-pair/) 拆值与下标；[1031](https://leetcode.com/problems/maximum-sum-of-two-non-overlapping-subarrays/) 两个不重叠定长窗口 |

例如 [1014](https://leetcode.com/problems/best-sightseeing-pair/) 的 `values[i]+values[j]+i-j` 拆为 `(values[i]+i)+(values[j]-j)`。扫描 j 时只需维护之前最大的 `values[i]+i`，查询后才能加入 j。[1031](https://leetcode.com/problems/maximum-sum-of-two-non-overlapping-subarrays/) 则把历史候选替换为已结束的左窗口，两个窗口的先后顺序各做一次。

PrefixTable 统一已完成状态的可见范围，不统一转移公式。[300](https://leetcode.com/problems/longest-increasing-subsequence/) 要取所有结尾最大值；[343](https://leetcode.com/problems/integer-break/) 必须真正拆分，不能直接返回 n。下面的源文件均已嵌入对应模板，可独立复制到 LeetCode。

## 复制与运行约定

每道题是一个独立代码组：复制下方列明的模板一次，再复制该题适配函数。题目 `.go` 文件已经包含相同模板，不依赖仓库内部导入，可直接作为单题提交。不同题目可能使用相同入口名（如 `rob`、`maxProfit`、`cherryPickup`），请分别复制，不能把全部适配器拼进同一个文件。以下代码使用 Go 1.21+ 的内置 `min/max`。

## 固定模板

### PrefixTable

<!-- template: PrefixTable -->
```go
// PrefixTable 的 dp[0] 是空前缀；step(i, dp[:i]) 只能读取已完成状态。
func PrefixTable[T any](n int, empty T, step func(i int, previous []T) T) []T {
	dp := make([]T, n+1)
	dp[0] = empty
	for i := 1; i <= n; i++ {
		dp[i] = step(i, dp[:i])
	}
	return dp
}
```

### Independent

<!-- template: Independent -->
```go
// Independent 维护旧前缀的“不选末项 / 选末项”收益，允许全部不选。
func Independent(nums []int) int {
	skip, take := 0, 0
	for _, value := range nums {
		skip, take = max(skip, take), skip+value
	}
	return max(skip, take)
}
```

### PrefixBest

<!-- template: PrefixBest -->
```go
// PrefixBest 先用严格更早的候选查询，再把当前候选加入历史最优集合。
func PrefixBest(n int, initial int, evaluate func(i, best int), candidate func(i int) int) {
	best := initial
	for i := 0; i < n; i++ {
		evaluate(i, best)
		best = max(best, candidate(i))
	}
}
```

## 逐题适配器

### [70](https://leetcode.com/problems/climbing-stairs/)

[代码](l70/70.go) · [测试](l70/70_test.go)

空阶梯有一种走法；每次来自前一或前两级。

需要模板：`PrefixTable`。

<!-- solution: dp/linear/l70/70.go; templates: PrefixTable -->
```go
// 空阶梯有一种走法；每次来自前一或前两级。
func climbStairs(n int) int {
	dp := PrefixTable(n, 1, func(i int, old []int) int {
		ways := old[i-1]
		if i >= 2 {
			ways += old[i-2]
		}
		return ways
	})
	return dp[n]
}
```

### [198](https://leetcode.com/problems/house-robber/)

[代码](l198/198.go) · [测试](l198/198_test.go)

相邻房屋不能同时选。

需要模板：`Independent`。

<!-- solution: dp/linear/l198/198.go; templates: Independent -->
```go
// 相邻房屋不能同时选。
func rob(nums []int) int { return Independent(nums) }
```

### [213](https://leetcode.com/problems/house-robber-ii/)

[代码](l213/213.go) · [测试](l213/213_test.go)

环拆成不选尾项与不选首项，取两者最大值。

需要模板：`Independent`。

<!-- solution: dp/linear/l213/213.go; templates: Independent -->
```go
// 环拆成不选尾项与不选首项，取两者最大值。
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(Independent(nums[:len(nums)-1]), Independent(nums[1:]))
}
```

### [300](https://leetcode.com/problems/longest-increasing-subsequence/)

[代码](l300/300.go) · [测试](l300/300_test.go)

dp[i] 是以 nums[i-1] 结尾的长度，答案要汇总全部结尾；O(n²)。

需要模板：`PrefixTable`。

<!-- solution: dp/linear/l300/300.go; templates: PrefixTable -->
```go
// dp[i] 是以 nums[i-1] 结尾的长度，答案要汇总全部结尾；O(n²)。
func lengthOfLIS(nums []int) int {
	answer := 0
	PrefixTable(len(nums), 0, func(i int, old []int) int {
		best := 1
		for j := 1; j < i; j++ {
			if nums[j-1] < nums[i-1] {
				best = max(best, old[j]+1)
			}
		}
		answer = max(answer, best)
		return best
	})
	return answer
}
```

### [343](https://leetcode.com/problems/integer-break/)

[代码](l343/343.go) · [测试](l343/343_test.go)

至少拆成两项；剩余部分可以不继续拆，也可以使用旧状态。

需要模板：`PrefixTable`。

<!-- solution: dp/linear/l343/343.go; templates: PrefixTable -->
```go
// 至少拆成两项；剩余部分可以不继续拆，也可以使用旧状态。
func integerBreak(n int) int {
	dp := PrefixTable(n, 0, func(i int, old []int) int {
		best := 0
		for first := 1; first < i; first++ {
			best = max(best, first*max(i-first, old[i-first]))
		}
		return best
	})
	return dp[n]
}
```

### [1014](https://leetcode.com/problems/best-sightseeing-pair/)

[代码](l1014/1014.go) · [测试](l1014/1014_test.go)

历史项是 values[i]+i；先查询再加入，确保 i<j。

需要模板：`PrefixBest`。

<!-- solution: dp/linear/l1014/1014.go; templates: PrefixBest -->
```go
// https://leetcode.com/problems/best-sightseeing-pair/
// 历史项是 values[i]+i；先查询再加入，确保 i<j。
func maxScoreSightseeingPair(values []int) int {
	if len(values) < 2 {
		return 0
	}
	answer := 0
	PrefixBest(len(values)-1, values[0], func(i, best int) {
		j := i + 1
		answer = max(answer, best+values[j]-j)
	}, func(i int) int { j := i + 1; return values[j] + j })
	return answer
}
```

### [1031](https://leetcode.com/problems/maximum-sum-of-two-non-overlapping-subarrays/)

[代码](l1031/1031.go) · [测试](l1031/1031_test.go)

枚举右窗口起点，历史最优左窗口保证不重叠；两种先后顺序都计算。

需要模板：`PrefixBest`。

<!-- solution: dp/linear/l1031/1031.go; templates: PrefixBest -->
```go
// https://leetcode.com/problems/maximum-sum-of-two-non-overlapping-subarrays/
// 枚举右窗口起点，历史最优左窗口保证不重叠；两种先后顺序都计算。
func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	prefix := make([]int, len(nums)+1)
	for i, value := range nums {
		prefix[i+1] = prefix[i] + value
	}
	ordered := func(leftLen, rightLen int) int {
		count := len(nums) - leftLen - rightLen + 1
		if count <= 0 {
			return 0
		}
		answer := 0
		initial := prefix[leftLen]
		PrefixBest(count, initial, func(i, best int) {
			rightStart := leftLen + i
			answer = max(answer, best+prefix[rightStart+rightLen]-prefix[rightStart])
		}, func(i int) int {
			end := leftLen + i + 1
			if end > len(nums)-rightLen {
				return 0
			}
			return prefix[end] - prefix[end-leftLen]
		})
		return answer
	}
	return max(ordered(firstLen, secondLen), ordered(secondLen, firstLen))
}
```

## 原始说明与其他材料

[改写前的完整分类笔记](legacy/README.before-template-rewrite.md)保存原说明、外部链接以及无扩展名片段的索引。上述适配器是当前实现；历史文字中的“尚未实现”等描述仅反映当时状态。原有测试保留，新增测试覆盖补全实现的关键边界。

```sh
go test ./dp/linear/...
```
