# 网格 DP：按依赖顺序分成六种子模板

[返回 DP 总目录](../README.md)

二维输入不代表只有二维状态。先画依赖边，再决定按格子、按行／步数，还是按 DAG 递归求值。

| 子模板 | 适用条件与不变量 | 题目 |
| --- | --- | --- |
| MonotoneGrid | 当前格只依赖已完成的上、左；边界是否合法由适配器判断 | [62](https://leetcode.com/problems/unique-paths/)、[63](https://leetcode.com/problems/unique-paths-ii/)、[64](https://leetcode.com/problems/minimum-path-sum/)；[1594](https://leetcode.com/problems/maximum-non-negative-product-in-a-matrix/) 的最大/最小乘积对 |
| ReverseGrid | 从终点倒推，依赖下、右已完成状态 | [174](https://leetcode.com/problems/dungeon-game/) 最低进入血量 |
| SquareSides | side[r+1][c+1] 是当前右下角的全一正方形最大边长 | [221](https://leetcode.com/problems/maximal-square/) 取最大再平方；[1277](https://leetcode.com/problems/count-square-submatrices-with-all-ones/) 累加边长 |
| Layers | 新层只能读取旧层，返回独立数组，不能覆盖旧层 | [120](https://leetcode.com/problems/triangle/)、[931](https://leetcode.com/problems/minimum-falling-path-sum/)、[1289](https://leetcode.com/problems/minimum-falling-path-sum-ii/)、[1937](https://leetcode.com/problems/maximum-number-of-points-with-cost/)；[741](https://leetcode.com/problems/cherry-pickup/)、[1463](https://leetcode.com/problems/cherry-pickup-ii/) 的双人状态 |
| MemoGrid | 状态依赖严格无环，缓存的是完整答案 | [329](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/) 严格递增、[2684](https://leetcode.com/problems/maximum-number-of-moves-in-a-grid/) 列号增加，以及双人问题的递归变体 |
| Filled | 为多位置状态建立独立行并填入正确不可达初值 | [741](https://leetcode.com/problems/cherry-pickup/)、[1463](https://leetcode.com/problems/cherry-pickup-ii/) 的分层适配器辅助函数 |

例如 [1594](https://leetcode.com/problems/maximum-non-negative-product-in-a-matrix/) 遇到负数时，原来最小的乘积可能变成最大的，因此每格必须保留一对极值，比较之前不能取模。[741](https://leetcode.com/problems/cherry-pickup/) 把往返看作两人同步前进，同一步的列号由步数与行号推出；[1463](https://leetcode.com/problems/cherry-pickup-ii/) 则直接用行号和两列位置。这些状态不能合并成普通二维路径和。

[1289](https://leetcode.com/problems/minimum-falling-path-sum-ii/) 维护不同列的最小、次小值，[1937](https://leetcode.com/problems/maximum-number-of-points-with-cost/) 用左右扫描求 `max(previous[k]-abs(k-c))`，二者都在 Layers 内保留 O(mn) 优化。所有模板遵循题目矩形输入约束；递归内存最坏由状态数和递归深度决定。

## 复制与运行约定

每道题是一个独立代码组：复制下方列明的模板一次，再复制该题适配函数。题目 `.go` 文件已经包含相同模板，不依赖仓库内部导入，可直接作为单题提交。不同题目可能使用相同入口名（如 `rob`、`maxProfit`、`cherryPickup`），请分别复制，不能把全部适配器拼进同一个文件。以下代码使用 Go 1.21+ 的内置 `min/max`。

## 固定模板

### MonotoneGrid

<!-- template: MonotoneGrid -->
```go
// MonotoneGrid 按行填表；cell 只能读取已完成的上、左前驱。
// 缺失前驱为零值，是否可用必须由 cell 根据 r/c 判断。
func MonotoneGrid[T any](rows, cols int, cell func(r, c int, up, left T) T) [][]T {
	dp := make([][]T, rows)
	for r := 0; r < rows; r++ {
		dp[r] = make([]T, cols)
		for c := 0; c < cols; c++ {
			var up, left T
			if r > 0 {
				up = dp[r-1][c]
			}
			if c > 0 {
				left = dp[r][c-1]
			}
			dp[r][c] = cell(r, c, up, left)
		}
	}
	return dp
}
```

### ReverseGrid

<!-- template: ReverseGrid -->
```go
// ReverseGrid 从右下倒推；cell 只能读取已完成的下、右状态。
func ReverseGrid[T any](rows, cols int, cell func(r, c int, down, right T) T) [][]T {
	dp := make([][]T, rows)
	for r := range dp {
		dp[r] = make([]T, cols)
	}
	for r := rows - 1; r >= 0; r-- {
		for c := cols - 1; c >= 0; c-- {
			var down, right T
			if r+1 < rows {
				down = dp[r+1][c]
			}
			if c+1 < cols {
				right = dp[r][c+1]
			}
			dp[r][c] = cell(r, c, down, right)
		}
	}
	return dp
}
```

### Layers

<!-- template: Layers -->
```go
// Layers 的 transition 读取上一层，返回独立的新一层；不能覆盖 previous 的切片内容。
func Layers[S any](count int, initial S, transition func(layer int, previous S) S) S {
	state := initial
	for layer := 0; layer < count; layer++ {
		state = transition(layer, state)
	}
	return state
}
```

### SquareSides

<!-- template: SquareSides -->
```go
// SquareSides[r+1][c+1] 表示以 (r,c) 为右下角的全一正方形最大边长。
func SquareSides(rows, cols int, isOne func(r, c int) bool) [][]int {
	side := make([][]int, rows+1)
	for r := range side {
		side[r] = make([]int, cols+1)
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if isOne(r, c) {
				side[r+1][c+1] = 1 + min(side[r][c+1], side[r+1][c], side[r][c])
			}
		}
	}
	return side
}
```

### Filled

<!-- template: Filled -->
```go
func Filled(rows, cols, value int) [][]int {
	result := make([][]int, rows)
	for r := range result {
		result[r] = make([]int, cols)
		for c := range result[r] {
			result[r][c] = value
		}
	}
	return result
}
```

### MemoGrid

<!-- template: MemoGrid -->
```go
// MemoGrid 仅适用于依赖无环的状态；缓存的零值也是合法答案。
func MemoGrid[S comparable, V any](evaluate func(S, func(S) V) V) func(S) V {
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

### [62](https://leetcode.com/problems/unique-paths/)

[代码](l62/62.go) · [测试](l62/62_test.go)

起点为 1，其他格累计上方和左方；边界缺失前驱贡献 0。

需要模板：`MonotoneGrid`。

<!-- solution: dp/grid/l62/62.go; templates: MonotoneGrid -->
```go
// 起点为 1，其他格累计上方和左方；边界缺失前驱贡献 0。
func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	dp := MonotoneGrid(m, n, func(r, c, up, left int) int {
		if r == 0 && c == 0 {
			return 1
		}
		return up + left
	})
	return dp[m-1][n-1]
}
```

### [63](https://leetcode.com/problems/unique-paths-ii/)

[代码](l63/63.go) · [测试](l63/63_test.go)

先处理障碍，避免把被挡起点初始化为 1。

需要模板：`MonotoneGrid`。

<!-- solution: dp/grid/l63/63.go; templates: MonotoneGrid -->
```go
// 先处理障碍，避免把被挡起点初始化为 1。
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := MonotoneGrid(m, n, func(r, c, up, left int) int {
		if obstacleGrid[r][c] == 1 {
			return 0
		}
		if r == 0 && c == 0 {
			return 1
		}
		return up + left
	})
	return dp[m-1][n-1]
}
```

### [64](https://leetcode.com/problems/minimum-path-sum/)

[代码](l64/64.go) · [测试](l64/64_test.go)

缺失前驱不能当成零代价捷径；第一行和第一列分别处理。

需要模板：`MonotoneGrid`。

<!-- solution: dp/grid/l64/64.go; templates: MonotoneGrid -->
```go
// 缺失前驱不能当成零代价捷径；第一行和第一列分别处理。
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := MonotoneGrid(m, n, func(r, c, up, left int) int {
		if r == 0 && c == 0 {
			return grid[r][c]
		}
		if r == 0 {
			return left + grid[r][c]
		}
		if c == 0 {
			return up + grid[r][c]
		}
		return min(up, left) + grid[r][c]
	})
	return dp[m-1][n-1]
}
```

### [174](https://leetcode.com/problems/dungeon-game/)

[代码](l174/174.go) · [测试](l174/174_test.go)

状态是进入当前格所需的最低血量，反向推导且始终至少为 1。

需要模板：`ReverseGrid`。

<!-- solution: dp/grid/l174/174.go; templates: ReverseGrid -->
```go
// 状态是进入当前格所需的最低血量，反向推导且始终至少为 1。
func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 1
	}
	m, n := len(dungeon), len(dungeon[0])
	dp := ReverseGrid(m, n, func(r, c, down, right int) int {
		if r == m-1 && c == n-1 {
			return max(1, 1-dungeon[r][c])
		}
		need := int(^uint(0) >> 1)
		if r+1 < m {
			need = min(need, down)
		}
		if c+1 < n {
			need = min(need, right)
		}
		return max(1, need-dungeon[r][c])
	})
	return dp[0][0]
}
```

### [120](https://leetcode.com/problems/triangle/)

[代码](l120/120.go) · [测试](l120/120_test.go)

每行从上一行相邻两个合法列转移；最终答案取末行最小值。

需要模板：`Layers`。

<!-- solution: dp/grid/l120/120.go; templates: Layers -->
```go
// 每行从上一行相邻两个合法列转移；最终答案取末行最小值。
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	last := Layers(len(triangle), []int(nil), func(r int, previous []int) []int {
		row := make([]int, len(triangle[r]))
		for c, value := range triangle[r] {
			if r == 0 {
				row[c] = value
				continue
			}
			best := int(^uint(0) >> 1)
			if c < len(previous) {
				best = min(best, previous[c])
			}
			if c > 0 {
				best = min(best, previous[c-1])
			}
			row[c] = best + value
		}
		return row
	})
	answer := last[0]
	for _, value := range last {
		answer = min(answer, value)
	}
	return answer
}
```

### [931](https://leetcode.com/problems/minimum-falling-path-sum/)

[代码](l931/931.go) · [测试](l931/931_test.go)

上一行同列及左右相邻列，读取独立旧层。

需要模板：`Layers`。

<!-- solution: dp/grid/l931/931.go; templates: Layers -->
```go
// 上一行同列及左右相邻列，读取独立旧层。
func minFallingPathSum(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	cols := len(matrix[0])
	last := Layers(len(matrix), []int(nil), func(r int, previous []int) []int {
		row := make([]int, cols)
		for c, value := range matrix[r] {
			if r == 0 {
				row[c] = value
				continue
			}
			best := previous[c]
			if c > 0 {
				best = min(best, previous[c-1])
			}
			if c+1 < cols {
				best = min(best, previous[c+1])
			}
			row[c] = best + value
		}
		return row
	})
	answer := last[0]
	for _, value := range last {
		answer = min(answer, value)
	}
	return answer
}
```

### [1289](https://leetcode.com/problems/minimum-falling-path-sum-ii/)

[代码](l1289/1289.go) · [测试](l1289/1289_test.go)

上一行必须不同列；维护来自不同下标的最小、次小值，重复最小值也算两个来源，O(mn)。

需要模板：`Layers`。

<!-- solution: dp/grid/l1289/1289.go; templates: Layers -->
```go
// 上一行必须不同列；维护来自不同下标的最小、次小值，重复最小值也算两个来源，O(mn)。
func minFallingPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	last := Layers(len(grid), []int(nil), func(r int, previous []int) []int {
		if r == 0 {
			return append([]int(nil), grid[0]...)
		}
		first, second, index := int(^uint(0)>>1), int(^uint(0)>>1), -1
		for c, value := range previous {
			if value < first {
				first, second, index = value, first, c
			} else if value < second {
				second = value
			}
		}
		row := make([]int, len(grid[r]))
		for c, value := range grid[r] {
			best := first
			if c == index {
				best = second
			}
			row[c] = best + value
		}
		return row
	})
	answer := last[0]
	for _, value := range last {
		answer = min(answer, value)
	}
	return answer
}
```

### [1937](https://leetcode.com/problems/maximum-number-of-points-with-cost/)

[代码](l1937/1937.go) · [测试](l1937/1937_test.go)

把 max(previous[k]-abs(k-c)) 拆为左右扫描，保留 O(mn) 优化与 int64 收益。

需要模板：`Layers`。

<!-- solution: dp/grid/l1937/1937.go; templates: Layers -->
```go
// 把 max(previous[k]-abs(k-c)) 拆为左右扫描，保留 O(mn) 优化与 int64 收益。
func maxPoints(points [][]int) int64 {
	if len(points) == 0 || len(points[0]) == 0 {
		return 0
	}
	last := Layers(len(points), []int64(nil), func(r int, previous []int64) []int64 {
		n := len(points[r])
		row := make([]int64, n)
		if r == 0 {
			for c, v := range points[r] {
				row[c] = int64(v)
			}
			return row
		}
		left, right := make([]int64, n), make([]int64, n)
		left[0], right[n-1] = previous[0], previous[n-1]
		for c := 1; c < n; c++ {
			left[c] = max(previous[c], left[c-1]-1)
		}
		for c := n - 2; c >= 0; c-- {
			right[c] = max(previous[c], right[c+1]-1)
		}
		for c, v := range points[r] {
			row[c] = int64(v) + max(left[c], right[c])
		}
		return row
	})
	answer := last[0]
	for _, v := range last {
		answer = max(answer, v)
	}
	return answer
}
```

### [221](https://leetcode.com/problems/maximal-square/)

[代码](l221/221.go) · [测试](l221/221_test.go)

状态是边长，取最大边长后平方。

需要模板：`SquareSides`。

<!-- solution: dp/grid/l221/221.go; templates: SquareSides -->
```go
// 状态是边长，取最大边长后平方。
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	side := SquareSides(len(matrix), len(matrix[0]), func(r, c int) bool { return matrix[r][c] == '1' })
	best := 0
	for _, row := range side {
		for _, value := range row {
			best = max(best, value)
		}
	}
	return best * best
}
```

### [1277](https://leetcode.com/problems/count-square-submatrices-with-all-ones/)

[代码](l1277/1277.go) · [测试](l1277/1277_test.go)

同一边长状态，累加每个格子可贡献的 1..side 个正方形。

需要模板：`SquareSides`。

<!-- solution: dp/grid/l1277/1277.go; templates: SquareSides -->
```go
// 同一边长状态，累加每个格子可贡献的 1..side 个正方形。
func countSquares(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	side := SquareSides(len(matrix), len(matrix[0]), func(r, c int) bool { return matrix[r][c] == 1 })
	answer := 0
	for _, row := range side {
		for _, value := range row {
			answer += value
		}
	}
	return answer
}
```

### [1594](https://leetcode.com/problems/maximum-non-negative-product-in-a-matrix/)

[代码](l1594/1594.go) · [测试](l1594/1594_test.go)

负数会交换极大与极小乘积；路径比较前不能取模，最后才对非负最大值取模。

需要模板：`MonotoneGrid`。

<!-- solution: dp/grid/l1594/1594.go; templates: MonotoneGrid -->
```go
// 负数会交换极大与极小乘积；路径比较前不能取模，最后才对非负最大值取模。
func maxProductPath(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}
	type product struct {
		low, high int64
		valid     bool
	}
	m, n := len(grid), len(grid[0])
	dp := MonotoneGrid(m, n, func(r, c int, up, left product) product {
		value := int64(grid[r][c])
		if r == 0 && c == 0 {
			return product{value, value, true}
		}
		result := product{}
		for _, source := range []product{up, left} {
			if !source.valid {
				continue
			}
			a, b := source.low*value, source.high*value
			if !result.valid {
				result = product{min(a, b), max(a, b), true}
			} else {
				result.low = min(result.low, a, b)
				result.high = max(result.high, a, b)
			}
		}
		return result
	})
	answer := dp[m-1][n-1].high
	if answer < 0 {
		return -1
	}
	return int(answer % 1000000007)
}
```

### [1463](https://leetcode.com/problems/cherry-pickup-ii/)

[代码](l1463/1463.go) · [测试](l1463/1463_test.go)

层是行号，状态是两列位置；同格奖励只计算一次；O(mn²)。

需要模板：`Layers`、`Filled`。

<!-- solution: dp/grid/l1463/1463.go; templates: Layers,Filled -->
```go
// 层是行号，状态是两列位置；同格奖励只计算一次；O(mn²)。
func cherryPickup(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	const invalid = -1 << 30
	initial := Filled(n, n, invalid)
	initial[0][n-1] = grid[0][0]
	if n > 1 {
		initial[0][n-1] += grid[0][n-1]
	}
	last := Layers(m-1, initial, func(layer int, previous [][]int) [][]int {
		r := layer + 1
		next := Filled(n, n, invalid)
		for c1 := 0; c1 < n; c1++ {
			for c2 := 0; c2 < n; c2++ {
				if previous[c1][c2] == invalid {
					continue
				}
				for d1 := -1; d1 <= 1; d1++ {
					for d2 := -1; d2 <= 1; d2++ {
						a, b := c1+d1, c2+d2
						if a < 0 || a >= n || b < 0 || b >= n {
							continue
						}
						reward := grid[r][a]
						if a != b {
							reward += grid[r][b]
						}
						next[a][b] = max(next[a][b], previous[c1][c2]+reward)
					}
				}
			}
		}
		return next
	})
	answer := 0
	for _, row := range last {
		for _, value := range row {
			answer = max(answer, value)
		}
	}
	return answer
}
```

### [741](https://leetcode.com/problems/cherry-pickup/)

[代码](l741/741.go) · [测试](l741/741_test.go)

把往返改为两人同步前进；步数推出列坐标，状态仅保存两行坐标，不可达状态不能参与加分。

需要模板：`Layers`、`Filled`。

<!-- solution: dp/grid/l741/741.go; templates: Layers,Filled -->
```go
// 把往返改为两人同步前进；步数推出列坐标，状态仅保存两行坐标，不可达状态不能参与加分。
func cherryPickup(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 || grid[0][0] == -1 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	const invalid = -1 << 30
	initial := Filled(m, m, invalid)
	initial[0][0] = grid[0][0]
	last := Layers(m+n-2, initial, func(layer int, previous [][]int) [][]int {
		step := layer + 1
		next := Filled(m, m, invalid)
		for r1 := 0; r1 < m; r1++ {
			for r2 := 0; r2 < m; r2++ {
				c1, c2 := step-r1, step-r2
				if c1 < 0 || c1 >= n || c2 < 0 || c2 >= n || grid[r1][c1] == -1 || grid[r2][c2] == -1 {
					continue
				}
				best := invalid
				for _, p1 := range []int{r1, r1 - 1} {
					for _, p2 := range []int{r2, r2 - 1} {
						if p1 >= 0 && p2 >= 0 {
							best = max(best, previous[p1][p2])
						}
					}
				}
				if best == invalid {
					continue
				}
				reward := grid[r1][c1]
				if r1 != r2 {
					reward += grid[r2][c2]
				}
				next[r1][r2] = best + reward
			}
		}
		return next
	})
	return max(0, last[m-1][m-1])
}
```

### [329](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/)

[代码](l329/329.go) · [测试](l329/329_test.go)

严格递增保证无环；每格答案是从该格出发的最长路径，O(mn)。

需要模板：`MemoGrid`。

<!-- solution: dp/grid/l329/329.go; templates: MemoGrid -->
```go
// 严格递增保证无环；每格答案是从该格出发的最长路径，O(mn)。
func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	solve := MemoGrid(func(s [2]int, next func([2]int) int) int {
		best := 1
		for _, d := range directions {
			r, c := s[0]+d[0], s[1]+d[1]
			if r >= 0 && r < m && c >= 0 && c < n && matrix[r][c] > matrix[s[0]][s[1]] {
				best = max(best, 1+next([2]int{r, c}))
			}
		}
		return best
	})
	answer := 0
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			answer = max(answer, solve([2]int{r, c}))
		}
	}
	return answer
}
```

### [2684](https://leetcode.com/problems/maximum-number-of-moves-in-a-grid/)

[代码](l2684/2684.go) · [测试](l2684/2684_test.go)

每步列号增加，返回移动边数而不是访问格数；从第一列所有起点取最大。

需要模板：`MemoGrid`。

<!-- solution: dp/grid/l2684/2684.go; templates: MemoGrid -->
```go
// 每步列号增加，返回移动边数而不是访问格数；从第一列所有起点取最大。
func maxMoves(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	solve := MemoGrid(func(s [2]int, next func([2]int) int) int {
		best := 0
		for d := -1; d <= 1; d++ {
			r, c := s[0]+d, s[1]+1
			if r >= 0 && r < m && c < n && grid[r][c] > grid[s[0]][s[1]] {
				best = max(best, 1+next([2]int{r, c}))
			}
		}
		return best
	})
	answer := 0
	for r := 0; r < m; r++ {
		answer = max(answer, solve([2]int{r, 0}))
	}
	return answer
}
```

### [1463](https://leetcode.com/problems/cherry-pickup-ii/) · recursion 变体

[代码](l1463/variants/recursion/1463.go) · [测试](l1463/variants/recursion/1463_test.go)

递归变体仍使用行号和两列坐标，保留独立记忆化方法。

需要模板：`MemoGrid`。

<!-- solution: dp/grid/l1463/variants/recursion/1463.go; templates: MemoGrid -->
```go
// 递归变体仍使用行号和两列坐标，保留独立记忆化方法。
func cherryPickup(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	const invalid = -1 << 30
	type state struct{ row, c1, c2 int }
	solve := MemoGrid(func(s state, next func(state) int) int {
		if s.c1 < 0 || s.c1 >= n || s.c2 < 0 || s.c2 >= n {
			return invalid
		}
		if s.row == m {
			return 0
		}
		reward := grid[s.row][s.c1]
		if s.c1 != s.c2 {
			reward += grid[s.row][s.c2]
		}
		best := invalid
		for a := -1; a <= 1; a++ {
			for b := -1; b <= 1; b++ {
				best = max(best, next(state{s.row + 1, s.c1 + a, s.c2 + b}))
			}
		}
		return reward + best
	})
	return solve(state{0, 0, n - 1})
}
```

### [741](https://leetcode.com/problems/cherry-pickup/) · recursion 变体

[代码](l741/variants/recursion/741.go) · [测试](l741/variants/recursion/741_test.go)

递归变体使用两个同步移动者的三维状态，非法状态不进入得分相加。

需要模板：`MemoGrid`。

<!-- solution: dp/grid/l741/variants/recursion/741.go; templates: MemoGrid -->
```go
// 递归变体使用两个同步移动者的三维状态，非法状态不进入得分相加。
func cherryPickup(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	const invalid = -1 << 30
	type state struct{ r1, c1, r2 int }
	solve := MemoGrid(func(s state, next func(state) int) int {
		c2 := s.r1 + s.c1 - s.r2
		if s.r1 < 0 || s.r1 >= m || s.c1 < 0 || s.c1 >= n || s.r2 < 0 || s.r2 >= m || c2 < 0 || c2 >= n {
			return invalid
		}
		if grid[s.r1][s.c1] == -1 || grid[s.r2][c2] == -1 {
			return invalid
		}
		if s.r1 == m-1 && s.c1 == n-1 {
			return grid[s.r1][s.c1]
		}
		best := max(next(state{s.r1 + 1, s.c1, s.r2 + 1}), next(state{s.r1 + 1, s.c1, s.r2}),
			next(state{s.r1, s.c1 + 1, s.r2 + 1}), next(state{s.r1, s.c1 + 1, s.r2}))
		if best == invalid {
			return invalid
		}
		reward := grid[s.r1][s.c1]
		if s.r1 != s.r2 {
			reward += grid[s.r2][c2]
		}
		return reward + best
	})
	return max(0, solve(state{0, 0, 0}))
}
```

## 原始说明与其他材料

[改写前的完整分类笔记](legacy/README.before-template-rewrite.md)保存原说明、外部链接以及无扩展名片段的索引。上述适配器是当前实现；历史文字中的“尚未实现”等描述仅反映当时状态。原有测试保留，新增测试覆盖补全实现的关键边界。

[391 几何材料](https://leetcode.com/problems/perfect-rectangle/) · [代码](legacy_geometry/l391/391.go)单独保留，矩形覆盖的面积／角点奇偶性不属于网格 DP，不能套用本页模板。

```sh
go test ./dp/grid/...
```
