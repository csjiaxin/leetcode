# 背包：容量方向与枚举对象共同决定含义

[返回 DP 总目录](../README.md)

先写清物品是否只能使用一次、目标是可达/数量/最值，以及是否区分选取顺序。

| 子模板 | 固定不变量 | 题目 |
| --- | --- | --- |
| Capacity(reuse=false) | 容量倒序，来源还没有使用当前物品；零重量每项也只更新一次 | 416、1049、494 |
| Capacity(reuse=true) | 物品在外层、容量正序，允许反复使用当前物品 | 322 最少个数、518 组合数 |
| TwoResources | 两个容量都倒序，每个输入字符串至多选一次 | 474 |
| MemoWays | 下标不减或剩余金额变小，缓存完整子问题计数 | 518 递归变体 |

例如 494 中 `P-N=target`、`P+N=total`，所以选择正号元素的总和必须为 `(total+target)/2`。先检查范围与奇偶，再用 0/1 方案计数。零元素会把计数翻倍，不能直接忽略。

518 的外层是硬币，统计组合；若改为外层金额、内层枚举最后一个数字，就变成 377 的有序序列计数，需要独立模板，不能只改正序或倒序。322 用 amount+1 表示不可达，读取不可达时不继续加一。

## 复制与运行约定

每道题是一个独立代码组：复制下方列明的模板一次，再复制该题适配函数。题目 `.go` 文件已经包含相同模板，不依赖仓库内部导入，可直接作为单题提交。不同题目可能使用相同入口名（如 `rob`、`maxProfit`、`cherryPickup`），请分别复制，不能把全部适配器拼进同一个文件。以下代码使用 Go 1.21+ 的内置 `min/max`。

## 固定模板

### Capacity

<!-- template: Capacity -->
```go
// Capacity 修改调用方提供的初始表；reuse=false 倒序，true 正序。
// 0/1 允许零重量（每项仍只处理一次）；可重复使用的重量必须为正。
func Capacity[T any](dp []T, weights []int, reuse bool, combine func(old, from T) T) []T {
	for _, weight := range weights {
		if reuse {
			for c := weight; c < len(dp); c++ {
				dp[c] = combine(dp[c], dp[c-weight])
			}
		} else {
			for c := len(dp) - 1; c >= weight; c-- {
				dp[c] = combine(dp[c], dp[c-weight])
			}
		}
	}
	return dp
}
```

### TwoResources

<!-- template: TwoResources -->
```go
// TwoResources：每项消耗两种资源，两个容量都倒序，每项只选一次。
func TwoResources(costs [][2]int, zeroLimit, oneLimit int) [][]int {
	dp := make([][]int, zeroLimit+1)
	for z := range dp {
		dp[z] = make([]int, oneLimit+1)
	}
	for _, cost := range costs {
		for z := zeroLimit; z >= cost[0]; z-- {
			for o := oneLimit; o >= cost[1]; o-- {
				dp[z][o] = max(dp[z][o], dp[z-cost[0]][o-cost[1]]+1)
			}
		}
	}
	return dp
}
```

### MemoWays

<!-- template: MemoWays -->
```go
// MemoWays 仅适用于依赖无环的状态；缓存的零值也是合法答案。
func MemoWays[S comparable, V any](evaluate func(S, func(S) V) V) func(S) V {
	memo := make(map[S]V)
	var solve func(S) V
	solve = func(s S) V {
		if value, ok := memo[s]; ok {
			return value
		}
		value := evaluate(s, solve)
		memo[s] = value
		return value
	}
	return solve
}
```

## 逐题适配器

### 416

[代码](l416/416.go) · [测试](l416/416_test.go)

0/1 可达性：每个元素最多用一次，查询总和的一半。

需要模板：`Capacity`。

<!-- solution: dp/knapsack/l416/416.go; templates: Capacity -->
```go
// https://leetcode.com/problems/partition-equal-subset-sum/
// 0/1 可达性：每个元素最多用一次，查询总和的一半。
func canPartition(nums []int) bool {
	total := 0
	for _, value := range nums {
		total += value
	}
	if total%2 != 0 {
		return false
	}
	dp := make([]bool, total/2+1)
	dp[0] = true
	Capacity(dp, nums, false, func(old, from bool) bool { return old || from })
	return dp[total/2]
}
```

### 1049

[代码](l1049/1049.go) · [测试](l1049/1049_test.go)

求不超过一半的最大可达子集和，剩余差值为 total-2*chosen。

需要模板：`Capacity`。

<!-- solution: dp/knapsack/l1049/1049.go; templates: Capacity -->
```go
// https://leetcode.com/problems/last-stone-weight-ii/
// 求不超过一半的最大可达子集和，剩余差值为 total-2*chosen。
func lastStoneWeightII(stones []int) int {
	total := 0
	for _, value := range stones {
		total += value
	}
	dp := make([]bool, total/2+1)
	dp[0] = true
	Capacity(dp, stones, false, func(old, from bool) bool { return old || from })
	for chosen := total / 2; chosen >= 0; chosen-- {
		if dp[chosen] {
			return total - 2*chosen
		}
	}
	return 0
}
```

### 494

[代码](l494/494.go) · [测试](l494/494_test.go)

P-N=target、P+N=total，所以 P=(total+target)/2；零元素使计数翻倍。

需要模板：`Capacity`。

<!-- solution: dp/knapsack/l494/494.go; templates: Capacity -->
```go
// https://leetcode.com/problems/target-sum/
// P-N=target、P+N=total，所以 P=(total+target)/2；零元素使计数翻倍。
func findTargetSumWays(nums []int, target int) int {
	total := 0
	for _, value := range nums {
		total += value
	}
	if target < -total || target > total || (total+target)%2 != 0 {
		return 0
	}
	chosen := (total + target) / 2
	dp := make([]int, chosen+1)
	dp[0] = 1
	Capacity(dp, nums, false, func(old, from int) int { return old + from })
	return dp[chosen]
}
```

### 322

[代码](l322/322.go) · [测试](l322/322_test.go)

完全背包求最少个数；不可达状态不能继续加一。

需要模板：`Capacity`。

<!-- solution: dp/knapsack/l322/322.go; templates: Capacity -->
```go
// https://leetcode.com/problems/coin-change/
// 完全背包求最少个数；不可达状态不能继续加一。
func coinChange(coins []int, amount int) int {
	inf := amount + 1
	dp := make([]int, amount+1)
	for c := 1; c <= amount; c++ {
		dp[c] = inf
	}
	Capacity(dp, coins, true, func(old, from int) int {
		if from == inf {
			return old
		}
		return min(old, from+1)
	})
	if dp[amount] == inf {
		return -1
	}
	return dp[amount]
}
```

### 518

[代码](l518/518.go) · [测试](l518/518_test.go)

硬币在外层、金额正序，统计组合而不是选币排列。

需要模板：`Capacity`。

<!-- solution: dp/knapsack/l518/518.go; templates: Capacity -->
```go
// https://leetcode.com/problems/coin-change-ii/
// 硬币在外层、金额正序，统计组合而不是选币排列。
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	Capacity(dp, coins, true, func(old, from int) int { return old + from })
	return dp[amount]
}
```

### 474

[代码](l474/474.go) · [测试](l474/474_test.go)

每个字符串消耗零和一两种资源，不是单容量背包。

需要模板：`TwoResources`。

<!-- solution: dp/knapsack/l474/474.go; templates: TwoResources -->
```go
// 每个字符串消耗零和一两种资源，不是单容量背包。
func findMaxForm(strs []string, m int, n int) int {
	costs := make([][2]int, len(strs))
	for i, s := range strs {
		for _, char := range s {
			if char == '0' {
				costs[i][0]++
			} else {
				costs[i][1]++
			}
		}
	}
	return TwoResources(costs, m, n)[m][n]
}
```

### 518 · 历史独立版本

[代码](l518/legacy/recursive/518.go)

递归版固定硬币下标，只在同一下标重复选币，避免排列重复计数。

需要模板：`MemoWays`。

<!-- solution: dp/knapsack/l518/legacy/recursive/518.go; templates: MemoWays -->
```go
// 递归版固定硬币下标，只在同一下标重复选币，避免排列重复计数。
func change(amount int, coins []int) int {
	type state struct{ index, remaining int }
	solve := MemoWays(func(s state, next func(state) int) int {
		if s.remaining == 0 {
			return 1
		}
		if s.index == len(coins) || s.remaining < 0 {
			return 0
		}
		return next(state{s.index + 1, s.remaining}) + next(state{s.index, s.remaining - coins[s.index]})
	})
	return solve(state{0, amount})
}
```

## 原始说明与其他材料

[改写前的完整分类笔记](legacy/README.before-template-rewrite.md)保存原说明、外部链接以及无扩展名片段的索引。上述适配器是当前实现；历史文字中的“尚未实现”等描述仅反映当时状态。原有测试保留，新增测试覆盖补全实现的关键边界。

```sh
go test ./dp/knapsack/...
```
