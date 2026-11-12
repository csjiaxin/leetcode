# Dijkstra：按当前最优标签扩展

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

每题代码块包含自己的堆类型和建图，独立可编译；辅助类型重复是为单题阅读方便，不要把不同题代码块直接合并。

<a id="lc743"></a>

## 2. [743. 网络延迟时间](https://leetcode.com/problems/network-delay-time/)

[练习](l743/743.go) · [测试](l743/743_test.go)

**题意**：从k发信号，求所有节点收到信号的最短时间最大值，有不可达为-1。边有向，编号1..n。

**状态**：dist源为0，其余INF；candidate=d+w，用最小堆。

```go
package template

import "container/heap"

type state struct {
	node     int
	distance int64
}
type minHeap []state

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(state)) }
func (h *minHeap) Pop() any          { a := *h; x := a[len(a)-1]; *h = a[:len(a)-1]; return x }

func networkDelayTime(times [][]int, n int, k int) int {
	type edge struct {
		to     int
		weight int64
	}
	graph := make([][]edge, n)
	for _, e := range times {
		graph[e[0]-1] = append(graph[e[0]-1], edge{e[1] - 1, int64(e[2])})
	}
	const inf int64 = 1<<63 - 1
	dist := make([]int64, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[k-1] = 0
	queue := &minHeap{{k - 1, 0}}
	heap.Init(queue)
	for queue.Len() > 0 {
		cur := heap.Pop(queue).(state)
		if cur.distance != dist[cur.node] {
			continue
		}
		for _, e := range graph[cur.node] {
			if e.weight > inf-cur.distance {
				continue
			}
			candidate := cur.distance + e.weight
			if candidate < dist[e.to] {
				dist[e.to] = candidate
				heap.Push(queue, state{e.to, candidate})
			}
		}
	}
	var answer int64
	for _, d := range dist {
		if d == inf {
			return -1
		}
		answer = max(answer, d)
	}
	return int(answer)
}
```

**例子**：1→2(10)、1→3(1)、3→2(1)，dist[2]从10改善为2；旧堆项10要跳过。

**复杂度**：懒删除版本时间O(V+E log(E+2))，空间O(V+E)。题设非负且有限距离远小于INF。

<a id="lc1631"></a>

## 3. [1631. 最小体力消耗路径](https://leetcode.com/problems/path-with-minimum-effort/)

[练习](l1631/1631.go) · [测试](l1631/1631_test.go)

**题意**：路径代价为沿途相邻高度差的最大值，最小化这个最大值。

**变化点**：candidate=max(已有代价,当前边高度差)，不是相加；仍最小堆，因为扩展只会不减。

```go
package template

import "container/heap"

type state struct {
	row, col int
	effort   int64
}
type minHeap []state

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].effort < h[j].effort }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(state)) }
func (h *minHeap) Pop() any          { a := *h; x := a[len(a)-1]; *h = a[:len(a)-1]; return x }

func minimumEffortPath(heights [][]int) int {
	rows, cols := len(heights), len(heights[0])
	const inf int64 = 1<<63 - 1
	dist := make([][]int64, rows)
	for r := range dist {
		dist[r] = make([]int64, cols)
		for c := range dist[r] {
			dist[r][c] = inf
		}
	}
	dist[0][0] = 0
	queue := &minHeap{{0, 0, 0}}
	heap.Init(queue)
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for queue.Len() > 0 {
		cur := heap.Pop(queue).(state)
		r, c := cur.row, cur.col
		if cur.effort != dist[r][c] {
			continue
		}
		if r == rows-1 && c == cols-1 {
			return int(cur.effort)
		}
		for _, d := range directions {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
				continue
			}
			delta := int64(heights[nr][nc]) - int64(heights[r][c])
			if delta < 0 {
				delta = -delta
			}
			candidate := max(cur.effort, delta)
			if candidate < dist[nr][nc] {
				dist[nr][nc] = candidate
				heap.Push(queue, state{nr, nc, candidate})
			}
		}
	}
	return -1
}
```

**例子**：[[1,100],[2,3]]沿1→2→3的最大差1，优于经过100。单格代价0。

**复杂度**：V=RC，四邻网格E=O(V)，时间O(V log(V+1))，空间O(V)。高度差int64覆盖题设范围。

<a id="lc1514"></a>

## 4. [1514. 概率最大的路径](https://leetcode.com/problems/path-with-maximum-probability/)

[练习](l1514/1514.go) · [测试](l1514/1514_test.go)

**题意**：无向边有成功概率，路径概率相乘，求起点到终点最大概率。

**变化点**：源为1、其他为0；candidate=prob×edgeProb；更大才更新，使用最大堆。概率范围[0,1]保证扩展不增加。

```go
package template

import "container/heap"

type state struct {
	node        int
	probability float64
}
type maxHeap []state

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i].probability > h[j].probability }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)        { *h = append(*h, x.(state)) }
func (h *maxHeap) Pop() any          { a := *h; x := a[len(a)-1]; *h = a[:len(a)-1]; return x }

func maxProbability(n int, edges [][]int, succProb []float64, startNode int, endNode int) float64 {
	type edge struct {
		to          int
		probability float64
	}
	graph := make([][]edge, n)
	for i, e := range edges {
		a, b := e[0], e[1]
		graph[a] = append(graph[a], edge{b, succProb[i]})
		graph[b] = append(graph[b], edge{a, succProb[i]})
	}
	best := make([]float64, n)
	best[startNode] = 1
	queue := &maxHeap{{startNode, 1}}
	heap.Init(queue)
	for queue.Len() > 0 {
		cur := heap.Pop(queue).(state)
		if cur.probability != best[cur.node] {
			continue
		}
		if cur.node == endNode {
			return cur.probability
		}
		for _, e := range graph[cur.node] {
			candidate := cur.probability * e.probability
			if candidate > best[e.to] {
				best[e.to] = candidate
				heap.Push(queue, state{e.to, candidate})
			}
		}
	}
	return 0
}
```

**例子**：0→1为0.5、1→2为0.5、0→2为0.2，间接概率0.25更大。零概率路径与不可达的答案都可为0。

**复杂度**：时间O(V+E log(E+2))，空间O(V+E)。严格改善避免概率1的环重复入堆。

<a id="pitfalls"></a>

## 5. 对比与易错点

| 题目 | 初值 | 松弛 | 堆 |
| --- | --- | --- | --- |
| 743 | 源0，其余INF | d+w | 最小 |
| 1631 | 源0，其余INF | max(d,高度差) | 最小 |
| 1514 | 源1，其余0 | p×边概率 | 最大 |

- 首次入堆不代表确定，出堆先跳过过期标签，再考虑提前返回。
- 只在严格改善时入堆，避免等值环重复扩展。
- 743有向，1514无向，建图不能混用。
- 普通加法版本不能处理负边；INF也不应与合法有限值混淆。
- 每题完整堆实现各自独立，Less决定优先级，不能只改松弛而忘记改堆方向。

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
