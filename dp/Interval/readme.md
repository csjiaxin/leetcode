# 区间 DP：统一填表顺序，分别证明边界与操作

[返回 DP 总目录](../README.md)

本页实现统一采用闭区间 `[l,r]`，由短到长填表；转移依赖严格更短的区间。空区间不存入表，适配器按题意给出 0 等单位值。

| 题目 | 状态与操作时机 | 单点初值 |
| --- | --- | --- |
| 312 | 清空原下标 l..r；固定外邻居后枚举最后戳破者 | 仅有一个气球时的三个数乘积 |
| 1547 | l..r 是待切的切点；枚举第一次切割，外侧坐标决定段长 | 一次切割的整段长度 |
| 1039 | 顶点 l..r 的三角剖分；枚举与边界边相连的第三顶点 | 0；两点也没有三角形 |
| 375 | 先对错误分支取最坏 max，再对猜测取最好 min | 单个候选无需付费，0 |
| 664 | 最少打印次数；相同字符可复用一次打印 | 1 |
| 1000 | 压到当前长度可达到的最少堆数，分割步长为 k-1 | 0 |

312 常见写法是补边界后的开区间；这里明确换算成原数组闭区间，因此左外邻居是 values[l]，右外邻居是 values[r+2]。改变区间约定必须同时调整边界与空区间处理，不能只改循环符号。

1000 中每次减少 k-1 堆，所以最终可合并为一堆的前提是 `(n-1)%(k-1)==0`；区间长度决定剩余堆数，只有能再合成一堆时才增加区间总和。

每格枚举分割点通常 O(n³) 时间、O(n²) 空间；1000 的合法分割步长可减少常数。字符串区间适配器见 [字符串系列](../strings/readme.md)，博弈区间适配器见 [game](../game/README.md)。

## 复制与运行约定

每道题是一个独立代码组：复制下方列明的模板一次，再复制该题适配函数。题目 `.go` 文件已经包含相同模板，不依赖仓库内部导入，可直接作为单题提交。不同题目可能使用相同入口名（如 `rob`、`maxProfit`、`cherryPickup`），请分别复制，不能把全部适配器拼进同一个文件。以下代码使用 Go 1.21+ 的内置 `min/max`。

## 固定模板

### IntervalTable

<!-- template: IntervalTable -->
```go
// IntervalTable 按闭区间长度递增；step 只能读取严格更短区间。
// base 初始化单点；空区间不在表中，由适配器自行处理。
func IntervalTable[T any](n int, base func(i int) T, step func(l, r int, dp [][]T) T) [][]T {
	dp := make([][]T, n)
	for i := range dp {
		dp[i] = make([]T, n)
		dp[i][i] = base(i)
	}
	for length := 2; length <= n; length++ {
		for l := 0; l+length <= n; l++ {
			r := l + length - 1
			dp[l][r] = step(l, r, dp)
		}
	}
	return dp
}
```

## 逐题适配器

### 312

[代码](l312/312.go) · [测试](l312/312_test.go)

dp[l][r] 清空原下标闭区间 l..r，外邻居固定为补边界数组 values[l] 与 values[r+2]；枚举最后戳破者。

需要模板：`IntervalTable`。

<!-- solution: dp/interval/l312/312.go; templates: IntervalTable -->
```go
// https://leetcode.com/problems/burst-balloons/
// dp[l][r] 清空原下标闭区间 l..r，外邻居固定为补边界数组 values[l] 与 values[r+2]；枚举最后戳破者。
func maxCoins(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	values := make([]int, n+2)
	values[0], values[n+1] = 1, 1
	copy(values[1:], nums)
	dp := IntervalTable(n, func(i int) int { return values[i] * values[i+1] * values[i+2] },
		func(l, r int, dp [][]int) int {
			best := 0
			for last := l; last <= r; last++ {
				score := values[l] * values[last+1] * values[r+2]
				if last > l {
					score += dp[l][last-1]
				}
				if last < r {
					score += dp[last+1][r]
				}
				best = max(best, score)
			}
			return best
		})
	return dp[0][n-1]
}
```

### 1547

[代码](l1547/1547.go) · [测试](l1547/1547_test.go)

状态是尚待切割的切点闭区间；枚举第一次切割，段长来自外侧固定端点，不修改 cuts。

需要模板：`IntervalTable`。

<!-- solution: dp/interval/l1547/1547.go; templates: IntervalTable -->
```go
// https://leetcode.com/problems/minimum-cost-to-cut-a-stick/
// 状态是尚待切割的切点闭区间；枚举第一次切割，段长来自外侧固定端点，不修改 cuts。
func minCost(n int, cuts []int) int {
	if len(cuts) == 0 {
		return 0
	}
	positions := append([]int{0}, cuts...)
	positions = append(positions, n)
	// 切点最多 100 个；插入排序避免模板需要额外导入。
	for i := 1; i < len(positions); i++ {
		for j := i; j > 0 && positions[j] < positions[j-1]; j-- {
			positions[j], positions[j-1] = positions[j-1], positions[j]
		}
	}
	count := len(cuts)
	dp := IntervalTable(count, func(i int) int { return positions[i+2] - positions[i] },
		func(l, r int, dp [][]int) int {
			best := int(^uint(0) >> 1)
			for first := l; first <= r; first++ {
				cost := positions[r+2] - positions[l]
				if first > l {
					cost += dp[l][first-1]
				}
				if first < r {
					cost += dp[first+1][r]
				}
				best = min(best, cost)
			}
			return best
		})
	return dp[0][count-1]
}
```

### 1039

[代码](l1039/1039.go) · [测试](l1039/1039_test.go)

顶点闭区间 l..r；边界边 (l,r) 对应三角形第三点 k，两个子区间共享 k。

需要模板：`IntervalTable`。

<!-- solution: dp/interval/l1039/1039.go; templates: IntervalTable -->
```go
// 顶点闭区间 l..r；边界边 (l,r) 对应三角形第三点 k，两个子区间共享 k。
func minScoreTriangulation(values []int) int {
	n := len(values)
	if n < 3 {
		return 0
	}
	dp := IntervalTable(n, func(int) int { return 0 }, func(l, r int, dp [][]int) int {
		if r-l < 2 {
			return 0
		}
		best := int(^uint(0) >> 1)
		for k := l + 1; k < r; k++ {
			best = min(best, dp[l][k]+dp[k][r]+values[l]*values[k]*values[r])
		}
		return best
	})
	return dp[0][n-1]
}
```

### 375

[代码](l375/375.go) · [测试](l375/375_test.go)

先对猜错后的分支取最坏 max，再对猜测值取最好 min；下标对应数值 guess+1。

需要模板：`IntervalTable`。

<!-- solution: dp/interval/l375/375.go; templates: IntervalTable -->
```go
// https://leetcode.com/problems/guess-number-higher-or-lower-ii/
// 先对猜错后的分支取最坏 max，再对猜测值取最好 min；下标对应数值 guess+1。
func getMoneyAmount(n int) int {
	if n <= 1 {
		return 0
	}
	dp := IntervalTable(n, func(int) int { return 0 }, func(l, r int, dp [][]int) int {
		best := int(^uint(0) >> 1)
		for guess := l; guess <= r; guess++ {
			left, right := 0, 0
			if guess > l {
				left = dp[l][guess-1]
			}
			if guess < r {
				right = dp[guess+1][r]
			}
			best = min(best, guess+1+max(left, right))
		}
		return best
	})
	return dp[0][n-1]
}
```

### 664

[代码](l664/664.go) · [测试](l664/664_test.go)

相同字符可以复用一次打印；把 s[l] 与某个 s[k] 的打印合并，不能只做普通加法切分。

需要模板：`IntervalTable`。

<!-- solution: dp/interval/l664/664.go; templates: IntervalTable -->
```go
// https://leetcode.com/problems/strange-printer/
// 相同字符可以复用一次打印；把 s[l] 与某个 s[k] 的打印合并，不能只做普通加法切分。
func strangePrinter(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := IntervalTable(len(s), func(int) int { return 1 }, func(l, r int, dp [][]int) int {
		best := 1 + dp[l+1][r]
		for k := l + 1; k <= r; k++ {
			if s[k] != s[l] {
				continue
			}
			middle := 0
			if k > l+1 {
				middle = dp[l+1][k-1]
			}
			best = min(best, middle+dp[k][r])
		}
		return best
	})
	return dp[0][len(s)-1]
}
```

### 1000

[代码](l1000/1000.go) · [测试](l1000/1000_test.go)

每个区间压到 (长度-1)%(k-1)+1 堆；合法分割步长 k-1，能压成一堆时才加合并代价。

需要模板：`IntervalTable`。

<!-- solution: dp/interval/l1000/1000.go; templates: IntervalTable -->
```go
// https://leetcode.com/problems/minimum-cost-to-merge-stones/
// 每个区间压到 (长度-1)%(k-1)+1 堆；合法分割步长 k-1，能压成一堆时才加合并代价。
func mergeStones(stones []int, k int) int {
	n := len(stones)
	if n <= 1 {
		return 0
	}
	if k <= 1 || (n-1)%(k-1) != 0 {
		return -1
	}
	prefix := make([]int, n+1)
	for i, value := range stones {
		prefix[i+1] = prefix[i] + value
	}
	dp := IntervalTable(n, func(int) int { return 0 }, func(l, r int, dp [][]int) int {
		best := int(^uint(0) >> 1)
		for split := l; split < r; split += k - 1 {
			best = min(best, dp[l][split]+dp[split+1][r])
		}
		if (r-l)%(k-1) == 0 {
			best += prefix[r+1] - prefix[l]
		}
		return best
	})
	return dp[0][n-1]
}
```

## 原始说明与其他材料

[改写前的完整分类笔记](legacy/README.before-template-rewrite.md)保存原说明、外部链接以及无扩展名片段的索引。上述适配器是当前实现；历史文字中的“尚未实现”等描述仅反映当时状态。原有测试保留，新增测试覆盖补全实现的关键边界。

```sh
go test ./dp/interval/...
```
