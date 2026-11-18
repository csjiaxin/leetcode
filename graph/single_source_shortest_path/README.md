# Dijkstra：按当前最优标签扩展

## 算法选择与实际文件

| 条件／目标 | 模板 | 本仓库代码与题目 |
| --- | --- | --- |
| 非负加法最短路 | 最小堆，候选为 d+w，弹出过期项就跳过 | [通用 Dijkstra](dijkstra/dijkstra.go)；[743 练习](https://leetcode.com/problems/network-delay-time/) · [代码](l743/743.go)与下文参考解答 |
| 最小化路径最大边 | 最小堆，候选为 max(d,w) | [1631 练习](https://leetcode.com/problems/path-with-minimum-effort/) · [代码](l1631/1631.go)；[旧实现](l1631/variants/legacy/1631.go) |
| 最大路径概率 | 最大堆，候选为 p×概率，边概率在 [0,1] | [1514 练习](https://leetcode.com/problems/path-with-maximum-probability/) · [代码](l1514/1514.go)；[原始片段](l1514/variants/legacy/1514) |
| 含负权边、检测源可达负环 | Bellman-Ford，最多 n-1 轮松弛后检查 | [独立通用示例](bellman_ford/bellman_ford.go) |
| 限制最多走多少条边 | 每轮读取上一轮快照，不能串用本轮新值 | [787 实现](https://leetcode.com/problems/cheapest-flights-within-k-stops/) · [代码](l787/787.go) |
| 权重只有 0/1 | 0-1 BFS；本库旧版本用免费边闭包＋分层 | [1368 原始片段](https://leetcode.com/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/) · [代码](l1368/1368) |
| 等权最少步数 | FIFO BFS | [BFS 模板族](../bfs/README.md) |

[743](https://leetcode.com/problems/network-delay-time/)、[1514](https://leetcode.com/problems/path-with-maximum-probability/)、[1631](https://leetcode.com/problems/path-with-minimum-effort/) 的 `.go` 主解均已使用下方 `BestLabels` 实现，每个文件包含完整的本地 helper、标签类型和堆，不依赖仓库内部包；现有测试原样保留。下文只声明一次公共 helper，再展示各题适配函数；旧变体继续保留供比较。通用 Dijkstra 和 Bellman-Ford 的 Edge 类型、演示 main 仍位于独立目录，各自编译，不强行混合。

### 边数限制为什么需要上一轮快照

```go
// dist[src]=0，其余为不可达；reach 也可以用独立 bool 数组表示。
for step := 0; step < maxEdges; step++ {
    previous := append([]int64(nil), dist...)
    for _, edge := range edges {
        if !reachable(previous[edge.from]) { continue }
        candidate := previous[edge.from] + edge.weight // 先检查溢出
        if candidate < dist[edge.to] { dist[edge.to] = candidate }
    }
}
```

这里是转移骨架，`reachable` 与不可达表示由调用方定义。第 step 轮只读取上一轮，才能保证新增一条边；原地读取更新值可能在一轮中穿过整条链。[787](https://leetcode.com/problems/cheapest-flights-within-k-stops/) 的 k 次中转最多允许 k+1 条边。

其他保留片段：[743 旧版本](https://leetcode.com/problems/network-delay-time/) · [代码](l743/variants/legacy/743)、[1976 最短路径计数](https://leetcode.com/problems/number-of-ways-to-arrive-at-destination/) · [本地](l1976/1976_undirect)、[2039 网络空闲](https://leetcode.com/problems/the-time-when-the-network-becomes-idle/) · [本地](l2039/2039_undirect)、[2359 功能图可达距离](https://leetcode.com/problems/find-closest-node-to-given-two-nodes/) · [本地](l2359/2359_directed)。它们按问题归档，不表示都采用 Dijkstra。

[返回总目录](../../algorithm_templates/README.md)

[扩展练习：迁移点与边界](#extended-practice)

## 目录

- [1. 核心思路](#core)
- [2. 743：非负加法最短路](#lc743)
- [3. 1631：最小化最大高度差](#lc1631)
- [4. 1514：最大路径概率](#lc1514)
- [5. 对比与易错点](#pitfalls)

<a id="core"></a>

## 1. 核心思路

标准Dijkstra适用于非负边权。堆保存候选标签，同节点可有多个旧条目，弹出后与当前最佳值比较，跳过过期项。

```text
源标签初始化 → 堆取最优 → 过期则跳过 → 松弛出边
仅严格改善时更新并入堆
```

**不变量**：弹出的非过期最优标签可以确定，因为沿路径继续扩展不会变得更优。加法最短路要求非负；minimax用max扩展；概率用[0,1]乘法并改成最大堆。不能任意替换公式而不证明性质。

### 通用标签模板与变化点

固定循环只负责维护最佳标签、按优先级出堆、跳过过期项，以及在严格改善时更新。每题提供起点及初值、`next(state,label,emit)` 和 `better(candidate,previous)`；`next` 为每个合法后继调用 `emit`，`better` 同时决定堆的优先级与松弛条件。

这里声明一次完整 helper，后面的题目代码块只包含适配函数。复制时，将公共代码块和选定题目的函数放在同一个文件中，保留一次 `package`、一次 `import "container/heap"`，以及一套 `BestLabels`／`pathLabel`／`labelQueue` 声明。各题主解已各自附上完整副本，可独立提交。泛型需要 Go 1.18 及以上；适配函数中的内置 `max` 需要 Go 1.21 及以上。

```go
package template

import "container/heap"

// BestLabels 要求扩展标签不会优于当前标签；better 同时定义优先级和严格改善。
func BestLabels[S comparable, V comparable](start S, initial V, next func(S, V, func(S, V)), better func(V, V) bool) map[S]V {
	best := map[S]V{start: initial}
	queue := &labelQueue[S, V]{better: better}
	heap.Push(queue, pathLabel[S, V]{start, initial})
	for queue.Len() > 0 {
		current := heap.Pop(queue).(pathLabel[S, V])
		if best[current.node] != current.value {
			continue
		}
		next(current.node, current.value, func(node S, value V) {
			previous, seen := best[node]
			if !seen || better(value, previous) {
				best[node] = value
				heap.Push(queue, pathLabel[S, V]{node, value})
			}
		})
	}
	return best
}

type pathLabel[S comparable, V comparable] struct {
	node  S
	value V
}
type labelQueue[S comparable, V comparable] struct {
	items  []pathLabel[S, V]
	better func(V, V) bool
}

func (q labelQueue[S, V]) Len() int           { return len(q.items) }
func (q labelQueue[S, V]) Less(i, j int) bool { return q.better(q.items[i].value, q.items[j].value) }
func (q labelQueue[S, V]) Swap(i, j int)      { q.items[i], q.items[j] = q.items[j], q.items[i] }
func (q *labelQueue[S, V]) Push(value any)    { q.items = append(q.items, value.(pathLabel[S, V])) }
func (q *labelQueue[S, V]) Pop() any {
	last := len(q.items) - 1
	value := q.items[last]
	q.items = q.items[:last]
	return value
}
```

`best` 中不存在表示尚未发现，无需把合法标签与 INF 哨兵混在一起；需要区分缺失与零标签时使用 map 的双返回值。`better` 必须定义一致的严格优先关系；沿边扩展不能使当前标签变得更优，而且对同一后继的扩展必须保持标签的优劣关系，才能安全丢弃被支配的旧标签。状态还须包含所有影响后续转移的信息。仅满足 Go 的 `comparable` 约束并不能证明这些算法前提。

加法最短路使用非负边和 `a < b`，minimax 使用 `max` 与 `a < b`，概率使用 [0,1] 内的有限值相乘与 `a > b`；浮点标签不能含 NaN。加法在调用 `emit` 前检查溢出。此版本搜索到堆耗尽，返回全部可达状态的最优标签；若改为单目标提前返回，必须在跳过过期项之后判断目标。

<a id="lc743"></a>

## 2. [743. 网络延迟时间](https://leetcode.com/problems/network-delay-time/)

[练习](l743/743.go) · [测试](l743/743_test.go)

**题意**：从k发信号，求所有节点收到信号的最短时间最大值，有不可达为-1。边有向，编号1..n。

**状态**：常见数组写法为dist源0、其余INF；这里的map只存已发现节点，源标签为0。candidate=d+w，用 `a < b` 同时指定最小堆和严格松弛。

```go
func networkDelayTime(times [][]int, n int, k int) int {
	type edge struct {
		to     int
		weight int64
	}
	graph := make([][]edge, n)
	for _, e := range times {
		graph[e[0]-1] = append(graph[e[0]-1], edge{e[1] - 1, int64(e[2])})
	}
	const limit = int64(1<<63 - 1)
	best := BestLabels(k-1, int64(0), func(u int, distance int64, emit func(int, int64)) {
		for _, e := range graph[u] {
			if e.weight <= limit-distance {
				emit(e.to, distance+e.weight)
			}
		}
	}, func(a, b int64) bool { return a < b })
	if len(best) != n {
		return -1
	}
	var answer int64
	for _, distance := range best {
		answer = max(answer, distance)
	}
	return int(answer)
}
```

**例子**：1→2(10)、1→3(1)、3→2(1)，dist[2]从10改善为2；旧堆项10要跳过。

**复杂度**：懒删除版本时间O(V+E log(E+2))，空间O(V+E)。题设非负且有限距离在int64范围内；用 `limit-distance` 检查加法溢出，map无需INF哨兵。

<a id="lc1631"></a>

## 3. [1631. 最小体力消耗路径](https://leetcode.com/problems/path-with-minimum-effort/)

[练习](l1631/1631.go) · [测试](l1631/1631_test.go)

**题意**：路径代价为沿途相邻高度差的最大值，最小化这个最大值。

**变化点**：candidate=max(已有代价,当前边高度差)，不是相加；仍最小堆，因为扩展只会不减。

```go
func minimumEffortPath(heights [][]int) int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return 0
	}
	rows, cols := len(heights), len(heights[0])
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	best := BestLabels([2]int{0, 0}, int64(0), func(at [2]int, effort int64, emit func([2]int, int64)) {
		for _, direction := range directions {
			r, c := at[0]+direction[0], at[1]+direction[1]
			if r < 0 || r >= rows || c < 0 || c >= cols {
				continue
			}
			delta := int64(heights[r][c]) - int64(heights[at[0]][at[1]])
			if delta < 0 {
				delta = -delta
			}
			emit([2]int{r, c}, max(effort, delta))
		}
	}, func(a, b int64) bool { return a < b })
	return int(best[[2]int{rows - 1, cols - 1}])
}
```

**例子**：[[1,100],[2,3]]沿1→2→3的最大差1，优于经过100。单格代价0。

**复杂度**：V=RC，四邻网格E=O(V)，时间O(V log(V+1))，空间O(V)。高度差int64覆盖题设范围。

<a id="lc1514"></a>

## 4. [1514. 概率最大的路径](https://leetcode.com/problems/path-with-maximum-probability/)

[练习](l1514/1514.go) · [测试](l1514/1514_test.go)

**题意**：无向边有成功概率，路径概率相乘，求起点到终点最大概率。

**变化点**：源为1，其他节点尚未发现；缺失的map结果读取为0，符合无路的答案约定。candidate=prob×edgeProb；用 `a > b` 表示更大才更新，同时指定最大堆。概率范围[0,1]保证扩展不增加。

```go
func maxProbability(n int, edges [][]int, succProb []float64, startNode int, endNode int) float64 {
	type edge struct {
		to          int
		probability float64
	}
	graph := make([][]edge, n)
	for i, e := range edges {
		graph[e[0]] = append(graph[e[0]], edge{e[1], succProb[i]})
		graph[e[1]] = append(graph[e[1]], edge{e[0], succProb[i]})
	}
	best := BestLabels(startNode, 1.0, func(u int, probability float64, emit func(int, float64)) {
		for _, e := range graph[u] {
			emit(e.to, probability*e.probability)
		}
	}, func(a, b float64) bool { return a > b })
	return best[endNode]
}
```

**例子**：0→1为0.5、1→2为0.5、0→2为0.2，间接概率0.25更大。零概率路径与不可达的答案都可为0。

**复杂度**：时间O(V+E log(E+2))，空间O(V+E)。严格改善避免概率1的环重复入堆。

<a id="pitfalls"></a>

## 5. 对比与易错点

| 题目 | 初值 | 松弛 | 堆 |
| --- | --- | --- | --- |
| [743](https://leetcode.com/problems/network-delay-time/) | 源0，其余未发现（数组版为INF） | d+w | 最小 |
| [1631](https://leetcode.com/problems/path-with-minimum-effort/) | 源0，其余未发现（数组版为INF） | max(d,高度差) | 最小 |
| [1514](https://leetcode.com/problems/path-with-maximum-probability/) | 源1；未发现节点查询结果为0 | p×边概率 | 最大 |

- 首次入堆不代表确定，出堆先跳过过期标签，再考虑提前返回。
- 只在严格改善时入堆，避免等值环重复扩展。
- [743](https://leetcode.com/problems/network-delay-time/)有向，[1514](https://leetcode.com/problems/path-with-maximum-probability/)无向，建图不能混用。
- 普通加法版本不能处理负边；INF也不应与合法有限值混淆。
- 公共helper中的 `better` 同时控制严格松弛和堆方向；各题主解都复制完整堆实现，独立提交时不依赖其他文件。

已有分类：[最短路](.)。

<a id="extended-practice"></a>

## 扩展练习：迁移点与边界

新增题目只补充学习路线与解法总结，不自动生成未请求的实现或测试。

| 题目（原题） | 如何迁移当前模板 | 必须确认的条件 / 边界 |
| --- | --- | --- |
| [1976. 到达目的地的方案数](https://leetcode.com/problems/number-of-ways-to-arrive-at-destination/) | 距离改善覆盖ways，相等则累加ways | 本题正权保证按距离处理计数依赖；计数取模 |
| [1786. 从第一个节点到最后一个节点的受限路径数](https://leetcode.com/problems/number-of-restricted-paths-from-first-to-last-node/) | 从终点Dijkstra求距离，再沿距离严格递减的DAG计数 | 受限路径不是只数最短路 |
| [778. 水位上升的泳池中游泳](https://leetcode.com/problems/swim-in-rising-water/) | 路径标签为沿途最高格高度，candidate=max(d,邻格高度) | 源标签为起点高度，不是0 |
| [1368. 使网格图至少有一条有效路径的最小代价](https://leetcode.com/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/) | 沿箭头代价0，否则1，Dijkstra可用 | 因仅0/1权，0-1 BFS可优化到线性 |
| [2045. 到达目的地的第二短时间](https://leetcode.com/problems/second-minimum-time-to-reach-destination/) | 每点保留两个严格不同到达时间，加上红绿灯等待 | 不是标准单dist；同一最短时间重复到达不算第二短 |
| [505. 迷宫II](https://leetcode.com/problems/the-maze-ii/) | 每次滚到墙前才停，边权为本次滚动格数，Dijkstra求最短距离 | 不是每次滚动算一步；目标必须为可停位置，原题可能需要会员 |

**复习方法**：先独立写出状态、合法性或松弛条件，再指出相较本篇基础题改变的那一行；若前提不满足，主动换算法，不为统一模板牺牲正确性。
