# BFS / 多源 BFS：等权最短路

[返回总目录](../../algorithm_templates/README.md)

[扩展练习：迁移点与边界](#extended-practice)

## 目录

- [1. 核心思路](#core)
- [2. 1091：八方向最短路径](#lc1091)
- [3. 994：同步腐烂传播](#lc994)
- [4. 542：到最近0的距离](#lc542)
- [5. 对比与易错点](#pitfalls)

<a id="core"></a>

## 1. 核心思路

每步代价相同，FIFO队列按距离非递减处理。**首次入队就标记访问**，首次发现的距离即最短距离。

```text
源点入队并标记 → 依次弹出 → 遍历合法未访问邻居 → 标记距离并入队
多源：所有源以距离0一起入队，而不是分别搜索
```

变化点是源、邻居、可通行条件和答案单位。加权图不能直接套普通BFS。以下各题代码独立，网格输入允许修改的实现会明确说明。

<a id="lc1091"></a>

## 2. [1091. 二进制矩阵中的最短路径](https://leetcode.com/problems/shortest-path-in-binary-matrix/)

[练习](l1091/1091.go) · [测试](l1091/1091_test.go)

**题意**：0可通行、1阻塞，可走八方向，求左上到右下路径上的格子数，无路为-1。

**状态**：dist=-1表示未发现，源距离为1，因为路径计格子而不是边。

```go
package template

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] != 0 || grid[n-1][n-1] != 0 {
		return -1
	}
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}
	queue := [][2]int{{0, 0}}
	dist[0][0] = 1
	for head := 0; head < len(queue); head++ {
		r, c := queue[head][0], queue[head][1]
		if r == n-1 && c == n-1 {
			return dist[r][c]
		}
		for dr := -1; dr <= 1; dr++ {
			for dc := -1; dc <= 1; dc++ {
				if dr == 0 && dc == 0 {
					continue
				}
				nr, nc := r+dr, c+dc
				if nr < 0 || nr >= n || nc < 0 || nc >= n {
					continue
				}
				if grid[nr][nc] != 0 || dist[nr][nc] != -1 {
					continue
				}
				dist[nr][nc] = dist[r][c] + 1
				queue = append(queue, [2]int{nr, nc})
			}
		}
	}
	return -1
}
```

**例子**：[[0,1],[1,0]]可直接对角走，答案2；单格[[0]]答案1。

**复杂度**：时间、空间O(n²)。不修改输入。

<a id="lc994"></a>

## 3. [994. 腐烂的橘子](https://leetcode.com/problems/rotting-oranges/)

[练习](l994/994.go) · [测试](l994/994_test.go)

**题意**：每分钟腐烂橘子感染四邻的新鲜橘子，求全部腐烂所需时间，无解-1，无新鲜0。

**状态**：所有腐烂橘子同时为源，记录fresh；每次处理固定一层，感染新格时立即改成2，兼作访问标记。

```go
package template

func orangesRotting(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	queue := [][2]int{}
	fresh := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, [2]int{r, c})
			}
			if grid[r][c] == 1 {
				fresh++
			}
		}
	}
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	minutes, head := 0, 0
	for head < len(queue) && fresh > 0 {
		end := len(queue)
		for head < end {
			r, c := queue[head][0], queue[head][1]
			head++
			for _, d := range directions {
				nr, nc := r+d[0], c+d[1]
				if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] != 1 {
					continue
				}
				grid[nr][nc] = 2
				fresh--
				queue = append(queue, [2]int{nr, nc})
			}
		}
		minutes++
	}
	if fresh > 0 {
		return -1
	}
	return minutes
}
```

**例子**：[[2,1,1,1,2]]，两端同时传播，中间最后在第2分钟腐烂。

**复杂度**：时间、队列空间O(RC)。本实现修改输入；测试应复制网格。

<a id="lc542"></a>

## 4. [542. 01矩阵](https://leetcode.com/problems/01-matrix/)

[练习](l542/542.go) · [测试](l542/542_test.go)

**题意**：为每个格子求到最近0的四方向距离，题设至少有一个0。

**状态**：所有0以距离0入队，其他格子为-1；多源BFS得到最近源距离，不需要对每个1单独搜索。

```go
package template

func updateMatrix(mat [][]int) [][]int {
	rows, cols := len(mat), len(mat[0])
	dist := make([][]int, rows)
	queue := [][2]int{}
	for r := 0; r < rows; r++ {
		dist[r] = make([]int, cols)
		for c := 0; c < cols; c++ {
			dist[r][c] = -1
			if mat[r][c] == 0 {
				dist[r][c] = 0
				queue = append(queue, [2]int{r, c})
			}
		}
	}
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for head := 0; head < len(queue); head++ {
		r, c := queue[head][0], queue[head][1]
		for _, d := range directions {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= rows || nc < 0 || nc >= cols || dist[nr][nc] != -1 {
				continue
			}
			dist[nr][nc] = dist[r][c] + 1
			queue = append(queue, [2]int{nr, nc})
		}
	}
	return dist
}
```

**例子**：[[0,1,1,1,0]]得到[[0,1,2,1,0]]，每格选择最近的0。

**复杂度**：时间、空间O(RC)。不修改输入。

<a id="pitfalls"></a>

## 5. 对比与易错点

| 题目 | 源点 | 邻居 | 距离单位 |
| --- | --- | --- | --- |
| 1091 | 左上角 | 八方向 | 格子数，源为1 |
| 994 | 所有腐烂橘子 | 四方向 | 分钟，源为0 |
| 542 | 所有0 | 四方向 | 边数，源为0 |

- 入队时标记，不是出队时标记，避免重复发现。
- 分层时提前固定end，新加入的格子属于下一层。
- 994没有新鲜橘子返回0，有新鲜但不可达返回-1。
- 额外资源如钥匙、剩余障碍消除次数需要进入状态，不能仅按坐标去重。
- 带权问题请看 [Dijkstra](../SingleSourceShortestPath/README.md)；已有分类：[graph](..)。

<a id="extended-practice"></a>

## 扩展练习：迁移点与边界

新增题目只补充学习路线与解法总结，不自动生成未请求的实现或测试。

| 题目（原题） | 如何迁移当前模板 | 必须确认的条件 / 边界 |
| --- | --- | --- |
| [127. 单词接龙](https://leetcode.com/problems/word-ladder/) | 单词为状态，每次改一个字母形成等权边，BFS最短转换链 | 答案计单词数，终点不在字典直接无解；可双向BFS |
| [752. 打开转盘锁](https://leetcode.com/problems/open-the-lock/) | 四位密码状态，每位加减1形成八个邻居 | 起点若禁用先返回-1；入队时标记 |
| [1162. 地图分析](https://leetcode.com/problems/as-far-from-land-as-possible/) | 所有陆地为多源，求海洋到最近陆地距离的最大值 | 全海或全陆返回-1 |
| [1926. 迷宫中离入口最近的出口](https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/) | 入口单源，首次到边界的非入口格即最近出口 | 入口本身即使在边界也不算出口 |
| [1293. 网格中的最短路径](https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/) | 状态为行列与剩余消除次数，不能仅按坐标visited | 同坐标更多剩余资源可支配较少资源，须证明剪枝 |
| [909. 蛇梯棋](https://leetcode.com/problems/snakes-and-ladders/) | 棋盘格编号为状态，掷骰子一次为一条等权边 | 蛇梯跳转只按题意一次，注意交替行坐标映射 |

**复习方法**：先独立写出状态、合法性或松弛条件，再指出相较本篇基础题改变的那一行；若前提不满足，主动换算法，不为统一模板牺牲正确性。
