# 网格 DFS：连通块与洪水填充（Flood Fill）

[返回 DFS 目录](../README.md)

八道题使用同一套递归四方向模板，各题主文件均内联完整的 FloodFill，没有仓库包导入或跨包薄封装。[flood_fill.go](flood_fill.go) 保留为参考实现；[参考模板测试](flood_fill_test.go)检查先标记后回调、重复起点、阻塞边界与对角不连通。原先放在 bfs/grid 的实现实际采用 DFS，因此归入本族。

下文只声明一次模板，逐题代码块调用它；组合模板与任一题适配即可使用。实际单题文件已包含两者，可直接单文件提交，同一个包只保留一份模板声明。另保留 [200 的旧写法](https://leetcode.com/problems/number-of-islands/) · [代码](l200/variants/legacy/200)，不会覆盖当前带测试的版本。

[130](https://leetcode.com/problems/surrounded-regions/) 保留本仓库原有的“返回新棋盘、输入只读”接口；原题要求原地修改时，可使用 [BFS 130 的原地版本](../../bfs/l130/130.go)。内联模板解决文件依赖，提交时仍需采用目标接口的修改语义。

## 目录

- [1. 核心模型与不变量](#model)
- [2. 共同模板：独立 visited，不修改输入](#template)
- [3. 八道题的状态与汇总](#problems)
- [4. 边界泛洪](#boundary)
- [5. 整块判定：不能短路遍历](#validation)
- [6. 编号合并](#labeling)
- [7. 易错点与复杂度](#pitfalls)
- [8. 练习顺序、代码与测试](#practice)

<a id="model"></a>

## 1. 核心模型与不变量

把格子看作节点，把上下左右的连接看作边。一次 DFS 处理起点所在的整个合法连通块。

```text
选择起点：扫描所有格子，或只扫描边界
    ↓
遇到未访问的合法格子
    ↓
启动 DFS，立即设置 visited
    ↓
向四个方向扩展，处理整个连通块
    ↓
汇总面积、合法性、编号等信息
```

**不变量：每个合法格子最多处理一次；一次完整泛洪后，起点所在的合法连通块全部被标记。**

| 对比 | 网格连通块 DFS | 树上路径 DFS＋回溯 |
| --- | --- | --- |
| 状态含义 | 哪些格子已经处理过 | 当前祖先路径包含什么 |
| 返回后 | visited 保留，不撤销 | 共享路径状态通常需要撤销 |
| 原因 | 后续起点不能重复处理同一连通块 | 兄弟分支不能继承另一条路径的信息 |

“是否回溯”取决于状态含义，不取决于是否递归。网格图可能有环，因此必须先标记，再扩展。

<a id="template"></a>

## 2. 共同模板：独立 visited，不修改输入

以下模板面向矩形网格。调用方分配与网格同形的 `visited`，同一轮扫描的所有起点共享它。`canEnter` 只判断格子类型；是否访问过由模板统一判断。

```go
func FloodFill(
    rows, cols, startR, startC int,
    visited [][]bool,
    canEnter func(r, c int) bool,
    visit func(r, c int),
) {
    var dfs func(int, int)
    dfs = func(r, c int) {
        if r < 0 || r >= rows || c < 0 || c >= cols {
            return
        }
        if visited[r][c] || !canEnter(r, c) {
            return
        }
        visited[r][c] = true
        visit(r, c)
        dfs(r-1, c)
        dfs(r+1, c)
        dfs(r, c-1)
        dfs(r, c+1)
    }
    dfs(startR, startC)
}
```

职责分离：**原图描述问题，visited 描述访问历史，visit 回调负责累计答案或写入独立结果。** 回调不再承担访问标记，不需要 `undo`。没有额外处理时传入 `func(r, c int) {}`。

例如 [695](https://leetcode.com/problems/max-area-of-island/) 的完整调用方式（需与上面的模板放在同一包）：

```go
func maxAreaOfIsland(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }
    rows, cols := len(grid), len(grid[0])
    visited := make([][]bool, rows)
    for r := range visited {
        visited[r] = make([]bool, cols)
    }
    best := 0
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] != 1 || visited[r][c] {
                continue
            }
            area := 0
            FloodFill(rows, cols, r, c, visited,
                func(r, c int) bool { return grid[r][c] == 1 },
                func(r, c int) { area++ },
            )
            if area > best {
                best = area
            }
        }
    }
    return best
}
```

八题都不修改输入。[733](https://leetcode.com/problems/flood-fill/) 返回逐行深拷贝的新图像；[130](https://leetcode.com/problems/surrounded-regions/) 也返回新棋盘，因此本仓库使用 **`solve(board [][]byte) [][]byte`**，不同于 LeetCode 原题的原地 `solve(board [][]byte)`。提交原题时需要适配其接口；这里优先保持非原地约定。[827](https://leetcode.com/problems/making-a-large-island/) 的编号写入独立 `labels` 矩阵。

<a id="problems"></a>

## 3. 八道题的状态与汇总

所有题都由模板先判断 `!visited[r][c]` 并标记，表中仅列出各题差异：

| 题目 | 起点选择 | canEnter | visit 与最终汇总 |
| --- | --- | --- | --- |
| [733](https://leetcode.com/problems/flood-fill/) 图像渲染 | 给定起点 | 原图颜色等于起点原色 | 在结果副本中写入新颜色 |
| [200](https://leetcode.com/problems/number-of-islands/) 岛屿数量 | 扫描全图 | 原图为字节 `'1'` | 回调为空；每启动一块，答案加一 |
| [695](https://leetcode.com/problems/max-area-of-island/) 最大面积 | 扫描全图 | 原图为整数 `1` | 累计本块面积，更新最大值 |
| [1020](https://leetcode.com/problems/number-of-enclaves/) 飞地数量 | 所有边界 | 原图为 `1` | 回调为空；最后数未访问陆地格子 |
| [1254](https://leetcode.com/problems/number-of-closed-islands/) 封闭岛屿 | 先边界、后全图 | 原图为 `0`，注意陆水反转 | 先标边界连通陆地，再数未访问连通块 |
| [130](https://leetcode.com/problems/surrounded-regions/) 被围绕的区域 | 所有边界 | 原图为字节 `'O'` | visited 表示安全区域；在副本中翻转未访问的 `'O'` |
| [1905](https://leetcode.com/problems/count-sub-islands/) 子岛屿 | 扫描 grid2 | grid2 为 `1` | 检查 grid1 对应格子，整块合法才加一 |
| [827](https://leetcode.com/problems/making-a-large-island/) 最大人工岛 | 先陆地、后水格 | 第一阶段原图为 `1` | 写入 labels 并累计面积；第二阶段合并不同编号 |

[733](https://leetcode.com/problems/flood-fill/) 即使新旧颜色相同，也不会循环：visited 独立于颜色。当前实现仍返回独立副本，修改返回值不会影响输入。

<a id="boundary"></a>

## 4. 边界泛洪

[1020](https://leetcode.com/problems/number-of-enclaves/)、[1254](https://leetcode.com/problems/number-of-closed-islands/)、[130](https://leetcode.com/problems/surrounded-regions/) 都反向考虑：**从边界可达的目标区域，不是封闭的内部区域。**

1. 从首末行、首末列启动泛洪，使用同一份 visited。
2. 保留这些访问标记，不改写原图。
3. 扫描全图处理剩余目标格子。

四个角或单行、单列的重复起点会被 visited 自动过滤。注意汇总单位：[1020](https://leetcode.com/problems/number-of-enclaves/) 数**格子**，[1254](https://leetcode.com/problems/number-of-closed-islands/) 数**连通块**，[130](https://leetcode.com/problems/surrounded-regions/) 在**结果副本**中翻转。[1254](https://leetcode.com/problems/number-of-closed-islands/) 第二轮继续使用同一访问表，防止一块被重复计数。

<a id="validation"></a>

## 5. 整块判定：不能短路遍历

[1905](https://leetcode.com/problems/count-sub-islands/) 中，即使发现不合法格子，仍需访问 grid2 的整个连通块。以下片段用于外层扫描发现新岛屿之后；visited 在扫描前统一分配：

```go
valid := true
FloodFill(rows, cols, r, c, visited,
    func(r, c int) bool { return grid2[r][c] == 1 },
    func(r, c int) {
        if grid1[r][c] != 1 {
            valid = false
        }
    },
)
if valid {
    count++
}
```

不要把 grid1 的条件放进 `canEnter`，否则会把不合法岛屿切碎；也不要因为 `valid == false` 就提前停止递归。两个输入网格都保持不变，访问历史只记在 visited 中。

<a id="labeling"></a>

## 6. 编号合并

[827](https://leetcode.com/problems/making-a-large-island/) 第一阶段仍使用模板：每发现一块陆地，分配正整数编号，把编号写入独立的 `labels[r][c]` 并累计 `areas[id]`。`labels == 0` 表示没有陆地编号，不需要保留原图中的 1，因此编号可从 1 开始。

第二阶段逐个检查原图中的水格，读取四邻居的 labels，候选面积为 **1＋相邻不同编号的面积之和**。每个水格使用自己的小集合去重：

- 必须按编号去重，不按面积去重。
- 同一块岛从多个方向接触水格，只加一次；不同岛面积相同，也都要加。
- 全水答案为 1；全陆答案为已有最大面积。
- 不必真的翻转候选格子，也不必为每个候选点重新 DFS。

<a id="pitfalls"></a>

## 7. 易错点与复杂度

- 四连通不是八连通，对角接触不相连。
- 先检查边界，再读 visited 或原图；先标记，再扩展。
- visited 在所有起点之间共享，递归返回不清除；不同函数调用重新分配。
- [200](https://leetcode.com/problems/number-of-islands/)、[130](https://leetcode.com/problems/surrounded-regions/) 使用字节 `'1'`、`'O'`，整数网格使用 `1`、`0`；[1254](https://leetcode.com/problems/number-of-closed-islands/) 的 0 是陆地。
- 不改变输入也不等于输出独立：[733](https://leetcode.com/problems/flood-fill/)、[130](https://leetcode.com/problems/surrounded-regions/) 必须逐行复制，不能只复制二维切片的外层。
- 接口按题目约束使用矩形网格；[1905](https://leetcode.com/problems/count-sub-islands/) 两个网格同形。空网格、零列均有保护。

时间均为 **O(mn)**，因为每个格子最多处理一次、每次只检查四邻居。[827](https://leetcode.com/problems/making-a-large-island/) 每个水格的编号集合最多四项，仍为 O(mn)。visited 为 **O(mn)**；递归栈最坏 O(mn)，不是 O(m+n)。[733](https://leetcode.com/problems/flood-fill/)、[130](https://leetcode.com/problems/surrounded-regions/) 的结果副本以及 [827](https://leetcode.com/problems/making-a-large-island/) 的 labels、面积表也至多 O(mn)。

大而深的连通块可改为显式栈 DFS 或队列 BFS，并在入栈/入队时标记。本族位于 `graph/dfs/grid`，主题是 DFS；两者都能完成相同的连通块填充，只是遍历顺序不同。

<a id="practice"></a>

## 8. 练习顺序、代码与测试

每个目录包含独立题解、相同的 FloodFill 模板及对应 `*_test.go`，可以单独阅读。所有测试都检查重复调用，计数类检查输入未变，返回图像/棋盘的题目还检查输出与输入不共享行。

| 顺序 | LeetCode | 本地题解 | 训练点 |
| --- | --- | --- | --- |
| 1 | [733. 图像渲染](https://leetcode.com/problems/flood-fill/) | [l733/733.go](l733/733.go) | 单起点同色填充 |
| 2 | [200. 岛屿数量](https://leetcode.com/problems/number-of-islands/) | [l200/200.go](l200/200.go) | 数连通块 |
| 3 | [695. 岛屿的最大面积](https://leetcode.com/problems/max-area-of-island/) | [l695/695.go](l695/695.go) | 累计块内面积 |
| 4 | [1020. 飞地的数量](https://leetcode.com/problems/number-of-enclaves/) | [l1020/1020.go](l1020/1020.go) | 边界泛洪后数格子 |
| 5 | [1254. 统计封闭岛屿的数目](https://leetcode.com/problems/number-of-closed-islands/) | [l1254/1254.go](l1254/1254.go) | 边界泛洪后数块，0 是陆地 |
| 6 | [130. 被围绕的区域](https://leetcode.com/problems/surrounded-regions/) | [l130/130.go](l130/130.go) | 标记安全区域，返回新棋盘 |
| 7 | [1905. 统计子岛屿](https://leetcode.com/problems/count-sub-islands/) | [l1905/1905.go](l1905/1905.go) | 整块判定，不短路 |
| 8 | [827. 最大人工岛](https://leetcode.com/problems/making-a-large-island/) | [l827/827.go](l827/827.go) | 额外编号矩阵，按编号合并 |

建议主线：**[200](https://leetcode.com/problems/number-of-islands/) → [695](https://leetcode.com/problems/max-area-of-island/) → [130](https://leetcode.com/problems/surrounded-regions/) → [1905](https://leetcode.com/problems/count-sub-islands/) → [827](https://leetcode.com/problems/making-a-large-island/)**，即数连通块 → 统计面积 → 边界泛洪 → 整块判定 → 编号合并。

```sh
go test -count=1 -cover ./graph/dfs/grid/...
```

测试覆盖空网格、单格、单行/单列、对角不连通、全水全陆、边界通道、同色填充、无效子岛屿与邻岛去重；[695](https://leetcode.com/problems/max-area-of-island/) 直接验证模板先标记、每格一次、重复起点无效。[827](https://leetcode.com/problems/making-a-large-island/) 对所有行列数均为 1～3 的二值网格做独立暴力 BFS 对照。

**一句话：原图只读，visited 永久记录访问历史；一次 Flood Fill 消化一个连通块，回调只负责汇总题目信息。**

## 9. 其余七题的完整模板适配

[695](https://leetcode.com/problems/max-area-of-island/) 的完整示例见第 2 节；下列函数继续只填写 canEnter、visit 和答案汇总，统一调用前面已声明的 FloodFill。文档组合模板与任一适配即可使用，实际单题主文件已经同时包含两者。

### [733](https://leetcode.com/problems/flood-fill/)：[本地完整文件](l733/733.go)

```go
// https://leetcode.com/problems/flood-fill/
// 返回逐行复制的新图像，即使新旧颜色相同也不与输入共享行。
func floodFill(image [][]int, sr, sc, color int) [][]int {
	if image == nil {
		return nil
	}
	result := make([][]int, len(image))
	for r := range image {
		if image[r] != nil {
			result[r] = append(make([]int, 0, len(image[r])), image[r]...)
		}
	}
	if len(image) == 0 || len(image[0]) == 0 {
		return result
	}
	rows, cols := len(image), len(image[0])
	if sr < 0 || sr >= rows || sc < 0 || sc >= cols {
		return result
	}
	visited := make([][]bool, rows)
	for r := range visited {
		visited[r] = make([]bool, cols)
	}
	original := image[sr][sc]
	FloodFill(rows, cols, sr, sc, visited,
		func(r, c int) bool { return image[r][c] == original },
		func(r, c int) { result[r][c] = color },
	)
	return result
}
```

### [200](https://leetcode.com/problems/number-of-islands/)：[本地完整文件](l200/200.go)

```go
// https://leetcode.com/problems/number-of-islands/
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for r := range visited {
		visited[r] = make([]bool, cols)
	}
	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != '1' || visited[r][c] {
				continue
			}
			count++
			FloodFill(rows, cols, r, c, visited,
				func(r, c int) bool { return grid[r][c] == '1' },
				func(r, c int) {},
			)
		}
	}
	return count
}
```

### [1020](https://leetcode.com/problems/number-of-enclaves/)：[本地完整文件](l1020/1020.go)

```go
// https://leetcode.com/problems/number-of-enclaves/
func numEnclaves(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for r := range visited {
		visited[r] = make([]bool, cols)
	}
	markBoundary := func(r, c int) {
		FloodFill(rows, cols, r, c, visited,
			func(r, c int) bool { return grid[r][c] == 1 },
			func(r, c int) {},
		)
	}
	for r := 0; r < rows; r++ {
		markBoundary(r, 0)
		markBoundary(r, cols-1)
	}
	for c := 0; c < cols; c++ {
		markBoundary(0, c)
		markBoundary(rows-1, c)
	}
	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			// 数剩余格子，不是连通块数。
			if grid[r][c] == 1 && !visited[r][c] {
				count++
			}
		}
	}
	return count
}
```

### [1254](https://leetcode.com/problems/number-of-closed-islands/)：[本地完整文件](l1254/1254.go)

```go
// https://leetcode.com/problems/number-of-closed-islands/
func closedIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for r := range visited {
		visited[r] = make([]bool, cols)
	}
	fill := func(r, c int) {
		FloodFill(rows, cols, r, c, visited,
			func(r, c int) bool { return grid[r][c] == 0 },
			func(r, c int) {},
		)
	}
	// 此题 0 是陆地，先标记与边界相连的陆地。
	for r := 0; r < rows; r++ {
		fill(r, 0)
		fill(r, cols-1)
	}
	for c := 0; c < cols; c++ {
		fill(0, c)
		fill(rows-1, c)
	}
	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 0 && !visited[r][c] {
				count++
				fill(r, c)
			}
		}
	}
	return count
}
```

### [130](https://leetcode.com/problems/surrounded-regions/)：[本地完整文件](l130/130.go)

```go
// https://leetcode.com/problems/surrounded-regions/
// 非原地版本：不同于 LeetCode 的原地 solve 接口，本函数返回新棋盘。
func solve(board [][]byte) [][]byte {
	if board == nil {
		return nil
	}
	result := make([][]byte, len(board))
	for r := range board {
		if board[r] != nil {
			result[r] = append(make([]byte, 0, len(board[r])), board[r]...)
		}
	}
	if len(board) == 0 || len(board[0]) == 0 {
		return result
	}
	rows, cols := len(board), len(board[0])
	visited := make([][]bool, rows)
	for r := range visited {
		visited[r] = make([]bool, cols)
	}
	markSafe := func(r, c int) {
		FloodFill(rows, cols, r, c, visited,
			func(r, c int) bool { return board[r][c] == 'O' },
			func(r, c int) {},
		)
	}
	for r := 0; r < rows; r++ {
		markSafe(r, 0)
		markSafe(r, cols-1)
	}
	for c := 0; c < cols; c++ {
		markSafe(0, c)
		markSafe(rows-1, c)
	}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == 'O' && !visited[r][c] {
				result[r][c] = 'X'
			}
		}
	}
	return result
}
```

### [1905](https://leetcode.com/problems/count-sub-islands/)：[本地完整文件](l1905/1905.go)

```go
// https://leetcode.com/problems/count-sub-islands/
// grid1、grid2 为同形矩形网格，二者均不修改。
func countSubIslands(grid1, grid2 [][]int) int {
	if len(grid2) == 0 || len(grid2[0]) == 0 {
		return 0
	}
	rows, cols := len(grid2), len(grid2[0])
	visited := make([][]bool, rows)
	for r := range visited {
		visited[r] = make([]bool, cols)
	}
	count := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid2[r][c] != 1 || visited[r][c] {
				continue
			}
			valid := true
			FloodFill(rows, cols, r, c, visited,
				func(r, c int) bool { return grid2[r][c] == 1 },
				func(r, c int) {
					// 只改变判定，不提前停止整块遍历。
					if grid1[r][c] != 1 {
						valid = false
					}
				},
			)
			if valid {
				count++
			}
		}
	}
	return count
}
```

### [827](https://leetcode.com/problems/making-a-large-island/)：[本地完整文件](l827/827.go)

```go
// https://leetcode.com/problems/making-a-large-island/
func largestIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	labels := make([][]int, rows)
	for r := range visited {
		visited[r] = make([]bool, cols)
		labels[r] = make([]int, cols)
	}
	areas := map[int]int{}
	id, best := 0, 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != 1 || visited[r][c] {
				continue
			}
			id++
			FloodFill(rows, cols, r, c, visited,
				func(r, c int) bool { return grid[r][c] == 1 },
				func(r, c int) {
					labels[r][c] = id
					areas[id]++
				},
			)
			if areas[id] > best {
				best = areas[id]
			}
		}
	}
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != 0 {
				continue
			}
			area := 1
			seen := map[int]bool{}
			for _, d := range directions {
				nr, nc := r+d[0], c+d[1]
				if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
					continue
				}
				neighborID := labels[nr][nc]
				// 按编号去重：同一块不能重复，不同块面积相同也都要加。
				if neighborID != 0 && !seen[neighborID] {
					seen[neighborID] = true
					area += areas[neighborID]
				}
			}
			if area > best {
				best = area
			}
		}
	}
	return best
}
```
