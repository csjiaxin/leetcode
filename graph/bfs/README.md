# BFS / 多源 BFS：等权最短路

## 归类与练习状态

本族只把所有边代价相同的问题当作最短路处理。单源、多源共享 FIFO 循环，区别在初始化；状态 BFS 必须把影响后续选择的资源／集合／历史加入去重键。

| 模板 | 固定不变量 | 题目 |
| --- | --- | --- |
| 单源最短路 | 首次入队确定最少步数 | 1091、1926、909、752、127、433、773、279、322 |
| 多源同时扩散 | 所有源距离为 0，一起入队 | 994、542、1162；934 先找到整座岛作为源 |
| 扩展状态 BFS | 只合并完整状态相同的搜索分支 | 1293 的剩余消除次数、864 的钥匙集合、847 的访问集合、1129 的边颜色、1654 的上一跳方向 |
| 其余建模与可达性 | 需要分别说明步数含义与重复扩展范围 | 130、1020、1311、1345、815、301 |

本目录24道题的 `.go` 文件均使用下方通用 BFS 辅助函数，每个主解文件都复制完整 helper，可独立提交，不依赖仓库内部包；已有测试保持原样。130、1020 的另一套已完成实现归入 [DFS 网格连通块](../dfs/grid/readme.md)，二者的接口与输入修改约定独立保留。[934 的 DFS＋BFS 旧笔记](l934/variants/dfs_bfs/934)也保留。

[返回仓库目录](../README.md)

[扩展练习：迁移点与边界](#extended-practice)

## 目录

- [1. 核心思路](#core)
  - [解决什么问题，为什么按层搜索？](#why-bfs)
  - [通用队列模板与四个变化点](#bfs-template)
  - [多源 BFS：多个起点同时扩散](#multi-source)
  - [状态 BFS：visited 到底标记什么？](#state-bfs)
  - [复杂度、算法边界与双向 BFS](#complexity)
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

变化点是源、邻居、可通行条件和答案单位。以下先解释原理，再分别给出题目实现；示例按各题合法输入约束编写，是否修改网格会明确说明。

<a id="why-bfs"></a>
### 解决什么问题，为什么按层搜索？

BFS 适合把问题表示为“状态＋一步转换”，且**每次转换代价相同**的场景：最少移动次数、最少操作次数、同步传播时间、到最近源点的距离。节点不一定是图上的整数，也可以是格子、密码、单词，甚至“格子＋钥匙集合”。

思路来自穷举，但按步数组织穷举：

```text
第0层：不走就能到达的状态，即所有起点
第1层：走1步能首次到达的状态
第2层：走2步能首次到达的状态
……
```

如果当前目标第一次出现在第 k 层，为什么不可能有更短的路径？因为所有少于 k 步的可达状态此前都已经被处理。FIFO 队列保证先发现的浅层状态先处理，后来追加的深层状态不能插队。

**不变量**：队列中的距离非递减；从距离 d 的状态扩展，只产生距离 d+1 的状态。于是第一次发现某个完整状态，就已得到它的最短距离。

DFS 先把一条路走深，因此“第一次找到”不一定最短；它可以搜索最短路，但通常需要更多枚举和额外比较。BFS 则利用等权条件，按层排除了所有更短的可能。

<a id="bfs-template"></a>
### 通用队列模板与四个变化点

先写清四项：**状态是什么 → 起点有哪些 → 一步如何生成邻居 → 何时完成**。随后填入同一个循环。

下面把固定的 FIFO 循环封装为泛型 `bfs`，各题只提供起点、邻居函数和可选的停止回调。`S` 必须可比较，坐标和扩展状态可用数组／结构体，字符串可直接作键；`dist` 同时保存距离和访问状态，整个搜索期间复用同一张表。显式邻接表示例中节点为 `0..len(graph)-1`，起点和目标均为合法节点；每条边代价为1。

```go
package template

// bfs visits each complete state once, at its minimum distance from any start.
// A true stop result ends the search before expanding that state; nil explores all.
func bfs[S comparable](starts []S, neighbors func(S) []S, stop func(S, int) bool) map[S]int {
	dist := make(map[S]int, len(starts))
	queue := make([]S, 0, len(starts))
	for _, start := range starts {
		if _, seen := dist[start]; seen {
			continue
		}
		dist[start] = 0
		queue = append(queue, start)
	}
	for head := 0; head < len(queue); head++ {
		current := queue[head]
		distance := dist[current]
		if stop != nil && stop(current, distance) {
			break
		}
		for _, next := range neighbors(current) {
			if _, seen := dist[next]; seen {
				continue
			}
			dist[next] = distance + 1
			queue = append(queue, next)
		}
	}
	return dist
}

func shortestDistance(graph [][]int, starts []int, target int) int {
	answer := -1
	bfs(starts, func(node int) []int {
		return graph[node]
	}, func(node, distance int) bool {
		if node == target {
			answer = distance
			return true
		}
		return false
	})
	return answer
}
```

| 固定部分 | 每道题改变的部分 |
| --- | --- |
| 队列先进先出、按 head 消费 | 状态类型 `S`，以及起点 `starts` |
| 首次发现时标记并入队 | 用数组／结构体／字符串表达完整的可比较状态 |
| 下一步距离为当前距离+1 | 枚举相邻格、旋转密码、修改字符等 |
| 处理到目标或队列耗尽 | `stop(state,distance)` 判断目标；传 `nil` 得到整张距离表 |

`neighbors(state)` 只返回合法的一步后继；helper 在入队前统一去重并赋值。`stop` 在出队后、生成邻居前执行，返回 `true` 立即停止；`nil` 会遍历全部可达状态。提前停止时返回的 `dist` 只包含已发现的状态。求单个目标的题目在回调中保存答案，避免搜索余下的状态。1311 只生成到指定层；301 收集第一个有合法括号串的整层，此后不再生成更深状态，遇到已入队的下一层立即结束。

README 中只声明一次 helper，下面的题解代码块与其放在同一包中即可运行；各练习主文件已各自附上完整副本，符合独立提交要求。Go 泛型需要 Go 1.18 及以上。

同样可以按层计数，替代为每个状态单独存距离：

```text
step = 0
所有起点入队，立即标记
while 队列非空:
    end = 当前队列长度          // 固定本层边界
    while head < end:
        current = 队列[head++]
        如果 current 完成目标，返回 step
        枚举合法未访问 next，标记并追加
    step++                     // 本层全部处理完才进入下一步
```

`end` 必须在处理这一层之前固定。若每次用不断增长的队列长度判断本层是否结束，就会把后续多层混在一起。Go 用切片加 `head` 即可，通常不需要链表；这种写法保留所有已入队元素，因此空间按全部发现状态计算。

<a id="multi-source"></a>
### 多源 BFS：多个起点同时扩散

把所有源点以距离0同时放入初始队列，后面循环完全不变：

```text
0 . . . 0       两端都是源
0 1 2 1 0       每个位置到最近源的距离
```

求的是 `dist[v] = min(从每个源到 v 的距离)`，不是分别跑完每个起点再合并，也不是把不同源的距离相加。

可以把它理解为“虚拟第0层包含所有源”。若用超级源帮助理解，它到各源应是零代价边；实际实现直接把各源设为0，不需要真的引入这些边。

- [994. 腐烂的橘子](https://leetcode.com/problems/rotting-oranges/)：所有腐烂橘子同时传播，最后感染时间是所需分钟数。
- [542. 01 矩阵](https://leetcode.com/problems/01-matrix/)：所有0同时扩展，每个格子第一次被发现时就遇到了最近的0。
- [847. 访问所有节点的最短路径](https://leetcode.com/problems/shortest-path-visiting-all-nodes/)：起点可选任意节点，因此初始化所有 `(i,1<<i)`，但它同时也是状态 BFS。

多源 BFS 能避免“每个点单独搜索最近源”造成的重复工作。

<a id="state-bfs"></a>
### 状态 BFS：visited 到底标记什么？

**状态必须包含所有会影响后续合法动作和目标判断的信息。** 判断两个状态能否合并，应问：它们未来可走的路、拥有的能力和完成条件是否等价？

| 问题 | 完整状态 | 为什么不能更少？ |
| --- | --- | --- |
| 普通静态迷宫 | `(row,col)` | 到同一格后可走方向一致 |
| 752 转盘锁 | 四位密码 | 不能只记录最后旋转的那一位 |
| 1293 消除障碍 | `(row,col,remain)` | 同一格，剩余消除次数影响后续通行能力 |
| 864 收集钥匙 | `(row,col,keyMask)` | 回到同一格时可能多了开锁能力 |
| 847 访问所有节点 | `(node,visitedMask)` | 同一节点，已访问集合不同，距离目标的进度不同 |
| 1129 交替颜色路径 | `(node,lastColor)` | 下一条可用边取决于上一条边的颜色 |

例如1293中先到某格时剩余0次，后到时剩余1次；后者可能穿过终点前的障碍，前者却不行。只用 `visited[row][col]` 会错误丢掉可行解。

最容易写对的版本：

```text
初始状态 (0,0,k) 入队，标记 visited[0][0][k]
弹出 (r,c,remain)
对每个相邻格 (nr,nc):
    如果越界：跳过
    nextRemain = remain - grid[nr][nc]
    如果 nextRemain<0 或该完整状态已访问：跳过
    标记 visited[nr][nc][nextRemain]，入队
```

先检查越界，再读取 `grid[nr][nc]`。本题起点为空地，每次移动算1步，单格起终点相同则答案为0。

进一步可记录每个格子已发现的最大剩余次数：**在标准 BFS 的非递减距离发现顺序下**，此前以不更多步数到达、且剩余次数不少的状态可以支配当前状态。先掌握完整状态版本，再证明和使用这种压缩；不能脱离距离条件直接比较资源。

位掩码是集合状态的紧凑表示：

```text
已访问集合 mask
访问节点 next：nextMask = mask | (1 << next)
所有节点都访问过：mask == (1 << n) - 1
```

这不是把不同状态强行合并，而是把集合编码为整数。847的源为 `(i,1<<i)`，访问过的节点允许再次经过；不能用普通“路径中禁止重复节点”的做法。

<a id="complexity"></a>
### 复杂度、算法边界与双向 BFS

复杂度应按**完整状态图**计算：显式图为 `O(V+E)` 时间、`O(V)` 空间；隐式图还需计算生成一个邻居的成本。

- R×C四/八方向网格：状态数 `RC`，每个状态常数个邻居，时间空间 `O(RC)`。
- 1293：最多 `RC(k+1)` 个状态，完整状态版本时间空间 `O(RC(k+1))`。
- 847：最多 `n·2^n` 个状态，时间上界 `O((n+E)·2^n)`，空间 `O(n·2^n)`；E为原图边数。
- 127：逐位置替换字符生成邻居时，还要算字符串构造和哈希成本，不能只按字典单词数估算时间。

| 条件 | 应考虑的方法 |
| --- | --- |
| 每条边代价相同且为正，求最少步数 | 普通 BFS，必要时把步数乘统一代价 |
| 边权只含0与1 | 0-1 BFS：双端队列，允许距离松弛 |
| 一般非负边权 | [Dijkstra](../single_source_shortest_path/README.md) |
| 求所有路径，而非最短距离 | 回溯/DFS；不能用全局 visited 随意合并路径 |
| 单一起终点且搜索分支大 | 可考虑双向 BFS；有向图从终点侧应走反向边 |

双向 BFS 通常每次扩展节点较少的一整层，减少探索量，但并不消除状态空间的最坏复杂度。不要只靠“两边碰到了”替代正确的层次、距离和去重管理。

**下列三题均调用同一个 `bfs`，对照的变化点只有状态、源、邻居与答案单位；其余练习的完整调用见各题 `.go` 文件。**

<a id="lc1091"></a>

## 2. [1091. 二进制矩阵中的最短路径](https://leetcode.com/problems/shortest-path-in-binary-matrix/)

[练习](l1091/1091.go) · [测试](l1091/1091_test.go)

**题意**：0可通行、1阻塞，可走八方向，求左上到右下路径上的格子数，无路为-1。

**状态**：格子坐标；helper 统一把源距离设为0，map 中不存在表示未发现。因为答案计格子数，在到达终点时返回边数加1。

```go
package template

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if n == 0 || len(grid[0]) == 0 || grid[0][0] != 0 || grid[n-1][n-1] != 0 {
		return -1
	}
	answer := -1
	bfs([][2]int{{0, 0}}, func(cell [2]int) [][2]int {
		next := make([][2]int, 0, 8)
		for dr := -1; dr <= 1; dr++ {
			for dc := -1; dc <= 1; dc++ {
				r, c := cell[0]+dr, cell[1]+dc
				if (dr != 0 || dc != 0) && r >= 0 && r < n && c >= 0 && c < n && grid[r][c] == 0 {
					next = append(next, [2]int{r, c})
				}
			}
		}
		return next
	}, func(cell [2]int, distance int) bool {
		if cell == [2]int{n-1, n-1} {
			answer = distance + 1 // This problem counts cells, including the start.
			return true
		}
		return false
	})
	return answer
}
```

**例子**：[[0,1],[1,0]]可直接对角走，答案2；单格[[0]]答案1。

**复杂度**：时间、空间O(n²)。不修改输入。

<a id="lc994"></a>

## 3. [994. 腐烂的橘子](https://leetcode.com/problems/rotting-oranges/)

[练习](l994/994.go) · [测试](l994/994_test.go)

**题意**：每分钟腐烂橘子感染四邻的新鲜橘子，求全部腐烂所需时间，无解-1，无新鲜0。

**状态**：所有腐烂橘子同时为源，记录fresh；helper 给出最少感染距离，最后出队的距离就是分钟数。生成新感染格时立即改成2并减少fresh。

```go
package template

func orangesRotting(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	starts := [][2]int{}
	fresh := 0
	for r := range grid {
		for c, orange := range grid[r] {
			if orange == 2 {
				starts = append(starts, [2]int{r, c})
			} else if orange == 1 {
				fresh++
			}
		}
	}
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	minutes := 0
	bfs(starts, func(cell [2]int) [][2]int {
		next := make([][2]int, 0, 4)
		for _, direction := range directions {
			r, c := cell[0]+direction[0], cell[1]+direction[1]
			if r >= 0 && r < rows && c >= 0 && c < cols && grid[r][c] == 1 {
				grid[r][c] = 2 // Mark when discovered, not when later removed.
				fresh--
				next = append(next, [2]int{r, c})
			}
		}
		return next
	}, func(_ [2]int, distance int) bool {
		minutes = distance
		return false
	})
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

**状态**：所有0以距离0入队，map 中不存在表示未发现；多源BFS得到最近源距离，最后按坐标填入结果矩阵，不需要对每个1单独搜索。

```go
package template

func updateMatrix(mat [][]int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return mat
	}
	rows, cols := len(mat), len(mat[0])
	starts := [][2]int{}
	for r := range mat {
		for c, value := range mat[r] {
			if value == 0 {
				starts = append(starts, [2]int{r, c})
			}
		}
	}
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	dist := bfs(starts, func(cell [2]int) [][2]int {
		next := make([][2]int, 0, 4)
		for _, direction := range directions {
			r, c := cell[0]+direction[0], cell[1]+direction[1]
			if r >= 0 && r < rows && c >= 0 && c < cols {
				next = append(next, [2]int{r, c})
			}
		}
		return next
	}, nil)
	answer := make([][]int, rows)
	for r := range answer {
		answer[r] = make([]int, cols)
		for c := range answer[r] {
			answer[r][c] = dist[[2]int{r, c}]
		}
	}
	return answer
}
```

**例子**：[[0,1,1,1,0]]得到[[0,1,2,1,0]]，每格选择最近的0。

**复杂度**：时间、空间O(RC)。不修改输入。

<a id="pitfalls"></a>

## 5. 对比与易错点

| 题目 | 源点 | 邻居 | 距离单位 |
| --- | --- | --- | --- |
| 1091 | 左上角 | 八方向 | 格子数，helper 源为0，结果加1 |
| 994 | 所有腐烂橘子 | 四方向 | 分钟，源为0 |
| 542 | 所有0 | 四方向 | 边数，源为0 |

- 入队时标记，不是出队时标记，避免重复发现。
- 分层时提前固定end，新加入的格子属于下一层。
- 994没有新鲜橘子返回0，有新鲜但不可达返回-1。
- 额外资源如钥匙、剩余障碍消除次数需要进入状态，不能仅按坐标去重。
- 带权问题请看 [Dijkstra](../single_source_shortest_path/README.md)；已有分类：[graph](../NOTES.md)。

<a id="extended-practice"></a>

## 扩展练习：迁移点与边界

共24道练习，按状态建模分类，而不只按题目的表面故事分类。每题都有 LeetCode 原题与本地目录链接。24个主解文件均已用同一 helper 实现；每个文件保留题目原函数签名和完整本地 helper，不导入仓库内部包。已有三题的测试保持不变，其他题目未新增测试。

### A. 基础网格与按层距离（4题）

| 原题 | 练习 | 状态、转移与完成条件 | 易错点 |
| --- | --- | --- | --- |
| [1091. 二进制矩阵中的最短路径](https://leetcode.com/problems/shortest-path-in-binary-matrix/) | [l1091](l1091/1091.go) | 格子为状态，八方向移动，到右下角结束 | 源距离为1；起终点阻塞直接无解 |
| [1926. 迷宫中离入口最近的出口](https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/) | [l1926](l1926/1926.go) | 入口为源，四方向，首次到非入口边界结束 | 入口即使在边界也不算出口；距离为边数 |
| [934. 最短的桥](https://leetcode.com/problems/shortest-bridge/) | [l934](l934/934.go) | 先找出一座岛，整座岛为多源，逐层扩展海水直到另一座岛 | 只统计翻转的海水格；先区分两座岛 |
| [909. 蛇梯棋](https://leetcode.com/problems/snakes-and-ladders/) | [l909](l909/909.go) | 编号为状态，一次骰子＋必要的一次跳转构成一条边 | 蛇梯不能连跳；先转换编号再读棋盘；访问标记记跳转后的格子 |

### B. 多源传播与边界可达性（6题）

| 原题 | 练习 | 状态、起点与扫描目标 | 易错点 |
| --- | --- | --- | --- |
| [994. 腐烂的橘子](https://leetcode.com/problems/rotting-oranges/) | [l994](l994/994.go) | 所有腐烂橘子为源，四方向感染，统计最后感染时间 | 无新鲜为0；无法全部感染为-1 |
| [542. 01 矩阵](https://leetcode.com/problems/01-matrix/) | [l542](l542/542.go) | 所有0为源，求每格最近源距离 | 不要对每个1各自搜索 |
| [1162. 地图分析](https://leetcode.com/problems/as-far-from-land-as-possible/) | [l1162](l1162/1162.go) | 所有陆地为源，海洋距离取最大 | 全海或全陆返回-1 |
| [1020. 飞地的数量](https://leetcode.com/problems/number-of-enclaves/) | [l1020](l1020/1020.go) | 所有边界陆地为源，标记能逃出的陆地，剩余陆地计数 | 这是可达性而非最短路；边界源要去重 |
| [130. 被围绕的区域](https://leetcode.com/problems/surrounded-regions/) | [l130](l130/130.go) | 从边界O标记安全区域，最后翻转未标记O | 本题原地修改，无返回值；不能先翻边界区域 |
| [1311. 获取你好友已观看的视频](https://leetcode.com/problems/get-watched-videos-by-your-friends/) | [l1311](l1311/1311.go) | 人为状态，单源BFS只取恰好level层，再统计视频 | 视频按频次、字典序排序；不同好友的相同视频需累计 |

### C. 隐式图：生成邻居就是建图（6题）

| 原题 | 练习 | 状态与一步转换 | 易错点 |
| --- | --- | --- | --- |
| [752. 打开转盘锁](https://leetcode.com/problems/open-the-lock/) | [l752](l752/752.go) | 四位密码为状态，每位循环加减1，共8个邻居 | 起点禁用优先判断；合法起点等于终点时答案0 |
| [127. 单词接龙](https://leetcode.com/problems/word-ladder/) | [l127](l127/127.go) | 单词为状态，每次只换一位，结果必须在字典内 | 答案计单词数；终点不在字典无解；可双向BFS |
| [433. 最小基因变化](https://leetcode.com/problems/minimum-genetic-mutation/) | [l433](l433/433.go) | 基因串为状态，每位尝试A/C/G/T且必须在基因库内 | 答案计变化次数，不像127计字符串个数 |
| [773. 滑动谜题](https://leetcode.com/problems/sliding-puzzle/) | [l773](l773/773.go) | 2×3棋盘编码为字符串，空位与合法邻位交换 | 预先列出每个空位下标的邻居；目标为123450 |
| [279. 完全平方数](https://leetcode.com/problems/perfect-squares/) | [l279](l279/279.go) | 和为状态，从0出发每步加一个平方数，目标n | 一步等于使用一个数；也可用DP，BFS用于建模训练 |
| [322. 零钱兑换](https://leetcode.com/problems/coin-change/) | [l322](l322/322.go) | 金额为状态，从0出发每步加一种硬币，目标amount | 硬币为正；amount=0返回0；这是DP题的BFS视角 |

### D. 扩展状态：资源、集合与历史限制（5题）

| 原题 | 练习 | 完整状态与转移 | 易错点 |
| --- | --- | --- | --- |
| [1293. 网格中的最短路径](https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/) | [l1293](l1293/1293.go) | `(r,c,remain)`，进入障碍格消耗一次资源 | 完整状态去重；起终点同格返回0 |
| [847. 访问所有节点的最短路径](https://leetcode.com/problems/shortest-path-visiting-all-nodes/) | [l847](l847/847.go) | `(node,mask)`，所有节点各自作为源，直到mask全1 | 可以重复经过节点；n=1返回0；状态数含2^n |
| [864. 获取所有钥匙的最短路径](https://leetcode.com/problems/shortest-path-to-get-all-keys/) | [l864](l864/864.go) | `(r,c,keyMask)`，拾取钥匙置位，过锁前查对应位 | 已拿钥匙后的重访是新状态；目标mask从实际钥匙集合计算 |
| [1129. 颜色交替的最短路径](https://leetcode.com/problems/shortest-path-with-alternating-colors/) | [l1129](l1129/1129.go) | `(node,lastColor)`，下一条边必须换色 | 起点应允许首条红边或蓝边；合并两种颜色状态的最短距离 |
| [1654. 到家的最少跳跃次数](https://leetcode.com/problems/minimum-jumps-to-reach-home/) | [l1654](l1654/1654.go) | `(position,lastWasBackward)`，可前跳，不能连续后跳 | 不越负位置、避禁区；需有界搜索，可用max(最大禁区位置+a,x)+b作为安全上界 |

### E. 图建模与去除重复扩展（3题）

| 原题 | 练习 | 如何建模和控制重复工作 | 易错点 |
| --- | --- | --- | --- |
| [815. 公交路线](https://leetcode.com/problems/bus-routes/) | [l815](l815/815.go) | 建站点→线路索引；按乘车次数分层，一次取出线路全部站点 | 线路只能展开一次，站点也要去重；source=target返回0 |
| [1345. 跳跃游戏 IV](https://leetcode.com/problems/jump-game-iv/) | [l1345](l1345/1345.go) | 下标为状态，邻居为相邻下标及同值下标；预建值→下标表 | 同值组展开后清空/标记，避免反复扫描退化到平方复杂度 |
| [301. 删除无效的括号](https://leetcode.com/problems/remove-invalid-parentheses/) | [l301](l301/301.go) | 字符串为状态，删一个括号是一条边，按删除次数分层 | 第一层有合法解时收集该层所有合法串，不再处理下一层；用集合去重 |

### 推荐刷题顺序

```text
建立队列与距离直觉：1091 → 1926
理解同时出发：994 → 542 → 1162
理解隐式建图：752 → 433 → 127 → 773
理解完整状态：1293 → 1129 → 864 → 847
学习搜索优化：1345 → 815 → 1654 → 301
```

如果普通 BFS 已熟，直接从 `542 → 752 → 1293 → 847` 开始。其他题用来复习：边界反向思考、层次聚合、DP与最短路之间的联系。

### 每题复盘清单

1. 一次操作是否真的代价相同？
2. 状态里是否包含所有影响未来的信息？
3. 初始队列是单源还是多源，helper 的源距离0如何换算成题目答案单位？
4. 邻居如何生成，何时判断合法，何时标记访问？
5. 求一个目标距离、整张距离表，还是某层所有答案？
6. 队列耗尽时该返回什么，起点本身满足目标时呢？
7. 复杂度按完整状态数与实际邻居生成成本计算了吗？

**练习提示**：复盘时重点检查起点即终点、不可达、重复到达、多源竞争和扩展状态。可用 `go test ./graph/bfs/...` 编译全部题目并运行已有测试；无测试目录编译通过不等于其所有边界都已获测试覆盖。

<a id="all-adapters"></a>

## 6. 其余 21 题的完整模板适配

前文给出 1091、994、542；下面补齐其余题目。每个代码块与第 1 节的同一份 bfs 组合，不重复声明队列遍历。1311 还需要标准库 sort；组合独立文件时，把导入放在 package 之后、所有声明之前。实际各题主文件已经包含所需导入及完整本地模板。

### 1020：[本地完整文件](l1020/1020.go)

```go
// https://leetcode.com/problems/number-of-enclaves/
func numEnclaves(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	starts := [][2]int{}
	land := 0
	for r := range grid {
		for c, value := range grid[r] {
			if value != 1 {
				continue
			}
			land++
			if r == 0 || r == rows-1 || c == 0 || c == cols-1 {
				starts = append(starts, [2]int{r, c})
			}
		}
	}
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	escaped := bfs(starts, func(cell [2]int) [][2]int {
		next := make([][2]int, 0, 4)
		for _, direction := range directions {
			r, c := cell[0]+direction[0], cell[1]+direction[1]
			if r >= 0 && r < rows && c >= 0 && c < cols && grid[r][c] == 1 {
				next = append(next, [2]int{r, c})
			}
		}
		return next
	}, nil)
	return land - len(escaped)
}
```

### 1129：[本地完整文件](l1129/1129.go)

```go
// https://leetcode.com/problems/shortest-path-with-alternating-colors/
func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	adjacent := [2][][]int{make([][]int, n), make([][]int, n)}
	for color, edges := range [][][]int{redEdges, blueEdges} {
		for _, edge := range edges {
			adjacent[color][edge[0]] = append(adjacent[color][edge[0]], edge[1])
		}
	}
	// Seeding both last colors allows either color for the first edge.
	dist := bfs([][2]int{{0, 0}, {0, 1}}, func(state [2]int) [][2]int {
		color := 1 - state[1]
		next := make([][2]int, 0, len(adjacent[color][state[0]]))
		for _, node := range adjacent[color][state[0]] {
			next = append(next, [2]int{node, color})
		}
		return next
	}, nil)
	answer := make([]int, n)
	for node := range answer {
		answer[node] = -1
		for color := 0; color < 2; color++ {
			if distance, found := dist[[2]int{node, color}]; found && (answer[node] == -1 || distance < answer[node]) {
				answer[node] = distance
			}
		}
	}
	return answer
}
```

### 1162：[本地完整文件](l1162/1162.go)

```go
// https://leetcode.com/problems/as-far-from-land-as-possible/
func maxDistance(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	starts := [][2]int{}
	for r := range grid {
		for c, value := range grid[r] {
			if value == 1 {
				starts = append(starts, [2]int{r, c})
			}
		}
	}
	if len(starts) == 0 || len(starts) == rows*cols {
		return -1
	}
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	answer := 0
	bfs(starts, func(cell [2]int) [][2]int {
		next := make([][2]int, 0, 4)
		for _, direction := range directions {
			r, c := cell[0]+direction[0], cell[1]+direction[1]
			if r >= 0 && r < rows && c >= 0 && c < cols {
				next = append(next, [2]int{r, c})
			}
		}
		return next
	}, func(_ [2]int, distance int) bool {
		answer = distance
		return false
	})
	return answer
}
```

### 127：[本地完整文件](l127/127.go)

```go
// https://leetcode.com/problems/word-ladder/
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 1
	}
	dictionary := make(map[string]bool, len(wordList))
	for _, word := range wordList {
		dictionary[word] = true
	}
	if !dictionary[endWord] {
		return 0
	}
	answer := 0
	bfs([]string{beginWord}, func(word string) []string {
		next := []string{}
		letters := []byte(word)
		for i, original := range letters {
			for replacement := byte('a'); replacement <= 'z'; replacement++ {
				if replacement == original {
					continue
				}
				letters[i] = replacement
				candidate := string(letters)
				if dictionary[candidate] {
					next = append(next, candidate)
				}
			}
			letters[i] = original
		}
		return next
	}, func(word string, distance int) bool {
		if word == endWord {
			answer = distance + 1 // Count words, rather than transformations.
			return true
		}
		return false
	})
	return answer
}
```

### 1293：[本地完整文件](l1293/1293.go)

```go
// https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/
func shortestPath(grid [][]int, k int) int {
	rows, cols := len(grid), len(grid[0])
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	answer := -1
	// The remaining budget belongs in the state; position alone is insufficient.
	bfs([][3]int{{0, 0, k}}, func(state [3]int) [][3]int {
		next := make([][3]int, 0, 4)
		for _, direction := range directions {
			r, c := state[0]+direction[0], state[1]+direction[1]
			if r < 0 || r >= rows || c < 0 || c >= cols {
				continue
			}
			remaining := state[2] - grid[r][c]
			if remaining >= 0 {
				next = append(next, [3]int{r, c, remaining})
			}
		}
		return next
	}, func(state [3]int, distance int) bool {
		if state[0] == rows-1 && state[1] == cols-1 {
			answer = distance
			return true
		}
		return false
	})
	return answer
}
```

### 130：[本地完整文件](l130/130.go)

```go
// https://leetcode.com/problems/surrounded-regions/
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	rows, cols := len(board), len(board[0])
	starts := [][2]int{}
	for r := range board {
		for c, value := range board[r] {
			if value == 'O' && (r == 0 || r == rows-1 || c == 0 || c == cols-1) {
				starts = append(starts, [2]int{r, c})
			}
		}
	}
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	safe := bfs(starts, func(cell [2]int) [][2]int {
		next := make([][2]int, 0, 4)
		for _, direction := range directions {
			r, c := cell[0]+direction[0], cell[1]+direction[1]
			if r >= 0 && r < rows && c >= 0 && c < cols && board[r][c] == 'O' {
				next = append(next, [2]int{r, c})
			}
		}
		return next
	}, nil)
	for r := range board {
		for c, value := range board[r] {
			if _, escaped := safe[[2]int{r, c}]; value == 'O' && !escaped {
				board[r][c] = 'X'
			}
		}
	}
}
```

### 1311：[本地完整文件](l1311/1311.go)

```go
import "sort"

// https://leetcode.com/problems/get-watched-videos-by-your-friends/
func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	frequency := map[string]int{}
	atLevel := false
	bfs([]int{id}, func(person int) []int {
		if atLevel {
			return nil // Every state at the requested depth is visited, never expanded.
		}
		return friends[person]
	}, func(person, distance int) bool {
		if distance == level {
			atLevel = true
			for _, title := range watchedVideos[person] {
				frequency[title]++
			}
		}
		return false
	})
	answer := make([]string, 0, len(frequency))
	for title := range frequency {
		answer = append(answer, title)
	}
	sort.Slice(answer, func(i, j int) bool {
		if frequency[answer[i]] != frequency[answer[j]] {
			return frequency[answer[i]] < frequency[answer[j]]
		}
		return answer[i] < answer[j]
	})
	return answer
}
```

### 1345：[本地完整文件](l1345/1345.go)

```go
// https://leetcode.com/problems/jump-game-iv/
func minJumps(arr []int) int {
	groups := map[int][]int{}
	for index, value := range arr {
		groups[value] = append(groups[value], index)
	}
	answer := -1
	bfs([]int{0}, func(index int) []int {
		next := make([]int, 0, len(groups[arr[index]])+2)
		if index > 0 {
			next = append(next, index-1)
		}
		if index+1 < len(arr) {
			next = append(next, index+1)
		}
		next = append(next, groups[arr[index]]...)
		// Its first expansion already offers every member at the earliest distance.
		delete(groups, arr[index])
		return next
	}, func(index, distance int) bool {
		if index == len(arr)-1 {
			answer = distance
			return true
		}
		return false
	})
	return answer
}
```

### 1654：[本地完整文件](l1654/1654.go)

```go
// https://leetcode.com/problems/minimum-jumps-to-reach-home/
func minimumJumps(forbidden []int, a int, b int, x int) int {
	blocked := make(map[int]bool, len(forbidden))
	maxForbidden := 0
	for _, position := range forbidden {
		blocked[position] = true
		if position > maxForbidden {
			maxForbidden = position
		}
	}
	// Beyond this bound, a detour can be shortened without crossing a forbidden point.
	upper := maxForbidden + a
	if x > upper {
		upper = x
	}
	upper += b
	if blocked[0] {
		return -1
	}
	answer := -1
	// The second field is 1 exactly when the last jump was backward.
	bfs([][2]int{{0, 0}}, func(state [2]int) [][2]int {
		next := make([][2]int, 0, 2)
		if forward := state[0] + a; forward <= upper && !blocked[forward] {
			next = append(next, [2]int{forward, 0})
		}
		if backward := state[0] - b; state[1] == 0 && backward >= 0 && !blocked[backward] {
			next = append(next, [2]int{backward, 1})
		}
		return next
	}, func(state [2]int, distance int) bool {
		if state[0] == x {
			answer = distance
			return true
		}
		return false
	})
	return answer
}
```

### 1926：[本地完整文件](l1926/1926.go)

```go
// https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/
func nearestExit(maze [][]byte, entrance []int) int {
	rows, cols := len(maze), len(maze[0])
	start := [2]int{entrance[0], entrance[1]}
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	answer := -1
	bfs([][2]int{start}, func(cell [2]int) [][2]int {
		next := make([][2]int, 0, 4)
		for _, direction := range directions {
			r, c := cell[0]+direction[0], cell[1]+direction[1]
			if r >= 0 && r < rows && c >= 0 && c < cols && maze[r][c] == '.' {
				next = append(next, [2]int{r, c})
			}
		}
		return next
	}, func(cell [2]int, distance int) bool {
		if cell != start && (cell[0] == 0 || cell[0] == rows-1 || cell[1] == 0 || cell[1] == cols-1) {
			answer = distance
			return true
		}
		return false
	})
	return answer
}
```

### 279：[本地完整文件](l279/279.go)

```go
// https://leetcode.com/problems/perfect-squares/
func numSquares(n int) int {
	squares := []int{}
	for value := 1; value*value <= n; value++ {
		squares = append(squares, value*value)
	}
	answer := -1
	bfs([]int{0}, func(sum int) []int {
		next := []int{}
		for _, square := range squares {
			if sum+square > n {
				break
			}
			next = append(next, sum+square)
		}
		return next
	}, func(sum, distance int) bool {
		if sum == n {
			answer = distance
			return true
		}
		return false
	})
	return answer
}
```

### 301：[本地完整文件](l301/301.go)

```go
// https://leetcode.com/problems/remove-invalid-parentheses/
func removeInvalidParentheses(s string) []string {
	valid := func(text string) bool {
		balance := 0
		for _, character := range text {
			if character == '(' {
				balance++
			} else if character == ')' {
				balance--
				if balance < 0 {
					return false
				}
			}
		}
		return balance == 0
	}
	answer := []string{}
	foundDepth := -1
	bfs([]string{s}, func(text string) []string {
		if foundDepth >= 0 {
			return nil // Keep checking this depth, but generate no deeper states.
		}
		next := []string{}
		for i := 0; i < len(text); i++ {
			if text[i] != '(' && text[i] != ')' {
				continue
			}
			if i > 0 && text[i] == text[i-1] {
				continue // Removing either identical adjacent bracket gives the same string.
			}
			next = append(next, text[:i]+text[i+1:])
		}
		return next
	}, func(text string, distance int) bool {
		if foundDepth >= 0 && distance > foundDepth {
			return true // Do not even validate the already queued deeper states.
		}
		if valid(text) {
			foundDepth = distance
			answer = append(answer, text)
		}
		return false
	})
	return answer
}
```

### 322：[本地完整文件](l322/322.go)

```go
// https://leetcode.com/problems/coin-change/
func coinChange(coins []int, amount int) int {
	answer := -1
	bfs([]int{0}, func(sum int) []int {
		next := make([]int, 0, len(coins))
		for _, coin := range coins {
			if coin > 0 && coin <= amount-sum {
				next = append(next, sum+coin)
			}
		}
		return next
	}, func(sum, distance int) bool {
		if sum == amount {
			answer = distance
			return true
		}
		return false
	})
	return answer
}
```

### 433：[本地完整文件](l433/433.go)

```go
// https://leetcode.com/problems/minimum-genetic-mutation/
func minMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}
	dictionary := make(map[string]bool, len(bank))
	for _, gene := range bank {
		dictionary[gene] = true
	}
	if !dictionary[endGene] {
		return -1
	}
	answer := -1
	bfs([]string{startGene}, func(gene string) []string {
		next := []string{}
		letters := []byte(gene)
		for i, original := range letters {
			for _, replacement := range []byte("ACGT") {
				if replacement == original {
					continue
				}
				letters[i] = replacement
				candidate := string(letters)
				if dictionary[candidate] {
					next = append(next, candidate)
				}
			}
			letters[i] = original
		}
		return next
	}, func(gene string, distance int) bool {
		if gene == endGene {
			answer = distance
			return true
		}
		return false
	})
	return answer
}
```

### 752：[本地完整文件](l752/752.go)

```go
// https://leetcode.com/problems/open-the-lock/
func openLock(deadends []string, target string) int {
	dead := make(map[string]bool, len(deadends))
	for _, state := range deadends {
		dead[state] = true
	}
	if dead["0000"] {
		return -1
	}
	answer := -1
	bfs([]string{"0000"}, func(state string) []string {
		next := make([]string, 0, 8)
		digits := []byte(state)
		for i, original := range digits {
			for _, change := range []int{1, 9} {
				digits[i] = byte((int(original-'0')+change)%10) + '0'
				candidate := string(digits)
				if !dead[candidate] {
					next = append(next, candidate)
				}
			}
			digits[i] = original
		}
		return next
	}, func(state string, distance int) bool {
		if state == target {
			answer = distance
			return true
		}
		return false
	})
	return answer
}
```

### 773：[本地完整文件](l773/773.go)

```go
// https://leetcode.com/problems/sliding-puzzle/
func slidingPuzzle(board [][]int) int {
	encoded := make([]byte, 0, 6)
	for _, row := range board {
		for _, tile := range row {
			encoded = append(encoded, byte(tile)+'0')
		}
	}
	adjacent := [][]int{{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}
	answer := -1
	bfs([]string{string(encoded)}, func(state string) []string {
		tiles := []byte(state)
		zero := 0
		for tiles[zero] != '0' {
			zero++
		}
		next := make([]string, 0, len(adjacent[zero]))
		for _, neighbor := range adjacent[zero] {
			tiles[zero], tiles[neighbor] = tiles[neighbor], tiles[zero]
			next = append(next, string(tiles))
			tiles[zero], tiles[neighbor] = tiles[neighbor], tiles[zero]
		}
		return next
	}, func(state string, distance int) bool {
		if state == "123450" {
			answer = distance
			return true
		}
		return false
	})
	return answer
}
```

### 815：[本地完整文件](l815/815.go)

```go
// https://leetcode.com/problems/bus-routes/
func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	stopRoutes := map[int][]int{}
	for route, stops := range routes {
		for _, stop := range stops {
			stopRoutes[stop] = append(stopRoutes[stop], route)
		}
	}
	expandedRoute := make([]bool, len(routes))
	answer := -1
	// One edge connects two stops reachable by boarding the same bus.
	bfs([]int{source}, func(stop int) []int {
		next := []int{}
		for _, route := range stopRoutes[stop] {
			if expandedRoute[route] {
				continue
			}
			expandedRoute[route] = true
			next = append(next, routes[route]...)
		}
		return next
	}, func(stop, distance int) bool {
		if stop == target {
			answer = distance
			return true
		}
		return false
	})
	return answer
}
```

### 847：[本地完整文件](l847/847.go)

```go
// https://leetcode.com/problems/shortest-path-visiting-all-nodes/
func shortestPathLength(graph [][]int) int {
	n := len(graph)
	starts := make([][2]int, n)
	for node := range graph {
		starts[node] = [2]int{node, 1 << node}
	}
	all := (1 << n) - 1
	answer := -1
	bfs(starts, func(state [2]int) [][2]int {
		next := make([][2]int, 0, len(graph[state[0]]))
		for _, node := range graph[state[0]] {
			next = append(next, [2]int{node, state[1] | (1 << node)})
		}
		return next
	}, func(state [2]int, distance int) bool {
		if state[1] == all {
			answer = distance
			return true
		}
		return false
	})
	return answer
}
```

### 864：[本地完整文件](l864/864.go)

```go
// https://leetcode.com/problems/shortest-path-to-get-all-keys/
func shortestPathAllKeys(grid []string) int {
	rows, cols := len(grid), len(grid[0])
	start := [3]int{}
	allKeys := 0
	for r, row := range grid {
		for c := range row {
			if row[c] == '@' {
				start = [3]int{r, c, 0}
			} else if row[c] >= 'a' && row[c] <= 'z' {
				allKeys |= 1 << (row[c] - 'a')
			}
		}
	}
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	answer := -1
	bfs([][3]int{start}, func(state [3]int) [][3]int {
		next := make([][3]int, 0, 4)
		for _, direction := range directions {
			r, c := state[0]+direction[0], state[1]+direction[1]
			if r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] == '#' {
				continue
			}
			cell, keys := grid[r][c], state[2]
			if cell >= 'A' && cell <= 'Z' && keys&(1<<(cell-'A')) == 0 {
				continue
			}
			if cell >= 'a' && cell <= 'z' {
				keys |= 1 << (cell - 'a')
			}
			next = append(next, [3]int{r, c, keys})
		}
		return next
	}, func(state [3]int, distance int) bool {
		if state[2] == allKeys {
			answer = distance
			return true
		}
		return false
	})
	return answer
}
```

### 909：[本地完整文件](l909/909.go)

```go
// https://leetcode.com/problems/snakes-and-ladders/
func snakesAndLadders(board [][]int) int {
	n := len(board)
	answer := -1
	bfs([]int{1}, func(square int) []int {
		next := make([]int, 0, 6)
		for roll := 1; roll <= 6 && square+roll <= n*n; roll++ {
			landing := square + roll
			rowFromBottom := (landing - 1) / n
			row, col := n-1-rowFromBottom, (landing-1)%n
			if rowFromBottom%2 == 1 {
				col = n - 1 - col
			}
			if board[row][col] != -1 {
				landing = board[row][col] // Exactly one snake/ladder per roll.
			}
			next = append(next, landing)
		}
		return next
	}, func(square, distance int) bool {
		if square == n*n {
			answer = distance
			return true
		}
		return false
	})
	return answer
}
```

### 934：[本地完整文件](l934/934.go)

```go
// https://leetcode.com/problems/shortest-bridge/
func shortestBridge(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	start := [2]int{}
	found := false
	for r := 0; r < rows && !found; r++ {
		for c, value := range grid[r] {
			if value == 1 {
				start, found = [2]int{r, c}, true
				break
			}
		}
	}
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	adjacent := func(cell [2]int) [][2]int {
		next := make([][2]int, 0, 4)
		for _, direction := range directions {
			r, c := cell[0]+direction[0], cell[1]+direction[1]
			if r >= 0 && r < rows && c >= 0 && c < cols {
				next = append(next, [2]int{r, c})
			}
		}
		return next
	}
	island := bfs([][2]int{start}, func(cell [2]int) [][2]int {
		next := [][2]int{}
		for _, neighbor := range adjacent(cell) {
			if grid[neighbor[0]][neighbor[1]] == 1 {
				next = append(next, neighbor)
			}
		}
		return next
	}, nil)
	starts := make([][2]int, 0, len(island))
	for cell := range island {
		starts = append(starts, cell)
	}
	answer := -1
	bfs(starts, adjacent, func(cell [2]int, distance int) bool {
		if _, firstIsland := island[cell]; grid[cell[0]][cell[1]] == 1 && !firstIsland {
			answer = distance - 1 // The final land cell is not a flipped sea cell.
			return true
		}
		return false
	})
	return answer
}
```
