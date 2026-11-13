# BFS / 多源 BFS：等权最短路

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

下面是显式邻接表图的可执行 Go 骨架，不强行封装回调。节点为 `0..len(graph)-1`，起点和目标均为合法节点；每条边代价为1。

```go
package template

func shortestDistance(graph [][]int, starts []int, target int) int {
    dist := make([]int, len(graph))
    for i := range dist { dist[i] = -1 }
    queue := []int{}
    for _, s := range starts {
        if dist[s] != -1 { continue } // 起点也要去重
        dist[s] = 0
        queue = append(queue, s)
    }
    for head := 0; head < len(queue); head++ {
        current := queue[head]
        if current == target { return dist[current] }
        for _, next := range graph[current] {
            if dist[next] != -1 { continue }
            dist[next] = dist[current] + 1 // 入队前标记
            queue = append(queue, next)
        }
    }
    return -1
}
```

| 固定部分 | 每道题改变的部分 |
| --- | --- |
| 队列先进先出、按 head 消费 | 队列元素的类型 |
| 首次发现时标记并入队 | 用数组、map，还是扩展状态表去重 |
| 下一步距离为当前距离+1 | 枚举相邻格、旋转密码、修改字符等 |
| 处理到目标或队列耗尽 | 目标是一个状态、全部感染，还是整张距离表 |

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
| 一般非负边权 | [Dijkstra](../graph/SingleSourceShortestPath/README.md) |
| 求所有路径，而非最短距离 | 回溯/DFS；不能用全局 visited 随意合并路径 |
| 单一起终点且搜索分支大 | 可考虑双向 BFS；有向图从终点侧应走反向边 |

双向 BFS 通常每次扩展节点较少的一整层，减少探索量，但并不消除状态空间的最坏复杂度。不要只靠“两边碰到了”替代正确的层次、距离和去重管理。

**本节模板负责理解通用流程；下列三题保留独立实现，便于对照状态、初始化和答案单位。**

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
- 带权问题请看 [Dijkstra](../graph/SingleSourceShortestPath/README.md)；已有分类：[graph](../graph/README.md)。

<a id="extended-practice"></a>

## 扩展练习：迁移点与边界

共24道练习，按状态建模分类，而不只按题目的表面故事分类。每题都有 LeetCode 原题与本地目录链接。新增目录只放**空函数声明**，沿用仓库的练习方式，不提前填写答案；带返回值的空函数尚不能编译，需独立完成后再测试。已有三题的代码与测试保持不变。

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
3. 初始队列是单源还是多源，源距离为0还是1？
4. 邻居如何生成，何时判断合法，何时标记访问？
5. 求一个目标距离、整张距离表，还是某层所有答案？
6. 队列耗尽时该返回什么，起点本身满足目标时呢？
7. 复杂度按完整状态数与实际邻居生成成本计算了吗？

**练习提示**：新题目录暂未添加测试，不要把“空函数无法编译”当成模板问题。先独立实现，再为起点即终点、不可达、重复到达、多源竞争和扩展状态补测试。已存在的三题可在实现完成后使用 `go test ./bfs/l1091 ./bfs/l994 ./bfs/l542`。
