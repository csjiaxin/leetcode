# 最短路的外层模板：Dijkstra 与 Bellman–Ford

## 目录

- [1. 外层调度与内层松弛](#layers)
- [2. Dijkstra：堆调度](#dijkstra)
- [3. 标准 Bellman–Ford：按轮扫描](#bellman-ford)
- [4. 限边 Bellman–Ford：上一轮快照](#bounded)
- [5. 对比与易错点](#comparison)
- [6. Bellman–Ford 与 Floyd–Warshall：共同的路径拼接](#ford-floyd)
  - [可复用的 RelaxVia](#relax-via)
  - [不同的外层循环](#ford-floyd-loops)
  - [初始化、负环与复杂度](#ford-floyd-summary)

<a id="layers"></a>

## 1. 外层调度与内层松弛

**外层模板负责调度：下一步处理谁、重复多少次、什么时候结束。**

**内层松弛负责计算：通过一条边，能否得到更好的答案。**

加法最短路沿边 `u → v` 的共同更新逻辑是：

```go
candidate := dist[u] + weight
if candidate < dist[v] {
    dist[v] = candidate
}
```

找到一条到达 v 的更短路径，就更新它的暂定最短距离，这叫松弛（relaxation）。更新并不一定成功，更不代表更新后立即得到最终最短距离。

通用模板可以将更新判断抽出来：

```go
relax := func(node S, candidate V) bool {
    previous, seen := best[node]
    if seen && !better(candidate, previous) {
        return false
    }
    best[node] = candidate
    return true
}
```

这里 `S`、`V`、`best`、`better` 由调用方提供；加法最短路使用 `better(a,b) = a < b`。`best` 未出现的节点表示尚未发现，不必把它与合法的零距离混淆。

两类算法共享 `best + relax` 的思想，但调度、正确性前提和终止条件应分别保留，不能只替换数据结构就认为算法等价。

以下数组骨架统一约定：顶点编号为 `0..vertexCount-1`，起点距离为 0，其余为 INF；只有从起点可达的节点才能扩展。INF 必须大于合法有限距离，距离加法需保证不溢出。代码片段不是单独可执行文件。

<a id="dijkstra"></a>

## 2. Dijkstra：堆调度

### 外层流程

```text
起点暂定最短距离设为 0，并入堆

堆不为空：
    弹出暂定最短距离最小的节点 u
    如果是过期条目，跳过
    枚举 u 的出边 u → v：
        计算候选距离
        如果松弛成功，把 v 和新距离入堆
```

```go
for pq.Len() > 0 {
    item := heap.Pop(pq).(Item)
    u, distance := item.node, item.dist
    if distance != dist[u] {
        continue
    }
    for _, edge := range graph[u] {
        v := edge.to
        candidate := distance + edge.weight
        if candidate < dist[v] {
            dist[v] = candidate
            heap.Push(pq, Item{node: v, dist: candidate})
        }
    }
}
```

`pq` 是按 dist 排序的最小堆，初始化时放入起点。堆中的距离是入堆时的快照；同一个节点可能有多个旧条目，所以弹出后需要与 `dist[u]` 比较。

### 不变量与前提

**每次取当前最近的节点：弹出的非过期暂定最短距离可以确定。** 标准加法最短路要求所有边权非负，因此绕经尚未处理的更远节点，不会再得到更短路径。

- 只有严格改善时入堆，避免零权环触发等值重复入堆。
- 首次入堆不代表答案确定。
- 单目标提前返回，必须放在过期检查之后。
- 负权边会破坏上述保证，不能直接套用。
- 这份标准模板不提供负环检测。

完整公共堆与通用实现见 [Dijkstra 笔记](README.md)。

<a id="bellman-ford"></a>

## 3. 标准 Bellman–Ford：按轮扫描

### 外层流程

```text
起点距离设为 0

最多执行 V-1 轮：
    扫描所有边，尝试从可达节点松弛
    整轮没有改善，则提前结束

再检查所有边：
    如果仍能改善，存在起点可达的负环
```

```go
for round := 0; round < vertexCount-1; round++ {
    changed := false
    for _, edge := range edges {
        u, v, weight := edge.from, edge.to, edge.weight
        if dist[u] == INF {
            continue
        }
        candidate := dist[u] + weight
        if candidate < dist[v] {
            dist[v] = candidate
            changed = true
        }
    }
    if !changed {
        break
    }
}
```

额外扫描检查负环：

```go
hasReachableNegativeCycle := false
for _, edge := range edges {
    if dist[edge.from] == INF {
        continue
    }
    if dist[edge.from]+edge.weight < dist[edge.to] {
        hasReachableNegativeCycle = true
        break
    }
}
```

### 为什么最多 V-1 轮？

没有起点可达的负环时，总能为每个可达节点选择一条不重复顶点的最短路径，最多包含 V-1 条边。每轮完整扫描至少能把最短距离沿路径再传播一条边，因此 V-1 轮足够。

**“至少传播一条”不等于“恰好传播一条”。** 上面的原地更新可以在同一轮连续利用刚更新的距离，具体取决于边的扫描顺序；这不影响标准 Bellman–Ford 的正确性，但不能据此严格限制路径边数。

可达负环存在时，受它影响的节点没有有限的最短距离。上述检查只判断存在性，不负责进一步标记所有受影响节点。起点不可达的负环不会被检测到；若要检测全图负环，需要额外的超级源等处理。

<a id="bounded"></a>

## 4. 限边 Bellman–Ford：上一轮快照

[787. K 站中转内最便宜的航班](https://leetcode.com/problems/cheapest-flights-within-k-stops/) 最多中转 K 次，也就是最多经过 K+1 条边。此时每轮必须严格只增加一条边的预算。

```go
for round := 0; round < maxEdges; round++ {
    next := append([]int(nil), dist...)
    for _, edge := range edges {
        if dist[edge.from] == INF {
            continue
        }
        candidate := dist[edge.from] + edge.weight
        if candidate < next[edge.to] {
            next[edge.to] = candidate
        }
    }
    dist = next
}
```

这里 `maxEdges = K+1`，循环开始前 dist 只记录起点距离 0，其余为 INF。

- 进入第 round 轮时，`dist` 表示最多使用 round 条边的最短距离。
- 本轮只读旧 `dist`，写新 `next`，结束后允许最多 round+1 条边。
- 先复制旧表，表示可以不增加边，保留“最多”而不是“恰好”的结果。

例如按顺序扫描 `0 → 1`、`1 → 2`，若原地更新，第一轮就可能到达 2，实际上用了两条边；使用快照时，第一轮最多只能从 0 到达 1。

有限边数预算下，即使有负环，答案也不会因为无限绕环而无下界；不能直接用“预算轮次后还能改善”作为标准负环判据。

<a id="comparison"></a>

## 5. 对比与易错点

| 对比 | Dijkstra | 标准 Bellman–Ford | 限边 Bellman–Ford |
| --- | --- | --- | --- |
| 下一步处理谁 | 堆取当前最近节点，扫描其出边 | 每轮扫描所有边 | 每轮扫描所有边 |
| 加法边权 | 必须非负 | 可以有负权 | 可以有负权 |
| 距离更新 | 原地，改善后入堆 | 可原地 | 读旧表、写新表 |
| 终止方式 | 堆空；单目标可在确定后提前结束 | 最多 V-1 轮；无改善可提前停 | 指定最多 L 轮 |
| 可达负环检测 | 当前模板不提供 | 再检查一轮能否改善 | 不能直接用有限预算后的改善判环 |
| 关键不变量 | 非过期最小距离出堆时确定 | 第 k 轮后，至少覆盖所有至多 k 条边路径的改善 | 第 k 轮后，严格表示至多 k 条边的最优结果 |

令 E 为边数，V 为节点数，L 为边数预算：标准 Bellman–Ford 时间 O(VE)，限边版本时间 O(LE)，距离辅助空间均为 O(V)。显式建图和存边的空间另计。二叉堆懒删除 Dijkstra 的时间可写为 O(V+E+(E+1)log(E+2))（数组初始化、建图和候选堆操作），堆最多保存 O(E+1) 条记录；简单图中通常简写为 O((V+E)log V)。

检查清单：

1. 先明确边权能否为负、是否限制边数，再选择外层调度。
2. 不可达节点不参与松弛；INF 运算不能溢出或伪造可达性。
3. Dijkstra 出堆先跳过旧条目，不能在首次发现目标时就返回。
4. 标准 Bellman–Ford 的 V-1 来自简单路径边数，不来自木棍长度、权重大小或其他数值范围。
5. 限边版本必须区分当前轮与上一轮，不能直接复用原地松弛循环。
6. 通用最优值模板用于最大概率、最小瓶颈等变体时，必须重新证明扩展的单调性，不能任意替换加法或比较符号。

**一句话：松弛决定怎样更新，外层模板决定按什么顺序、重复多少次更新。**

<a id="ford-floyd"></a>

## 6. Bellman–Ford 与 Floyd–Warshall：共同的路径拼接

这里的 Ford 指 Bellman–Ford。两者都可以理解为动态规划：**把两段已知路径拼接成候选路径，再与当前答案取最小值**。区别主要在状态维度和外层扩展的范围，不在松弛的本质。

```text
Bellman–Ford：起点 ──已知路径──→ u ──一条边──→ v
候选代价 = dist[u] + weight(u,v)

Floyd–Warshall：i ──已知路径──→ k ──已知路径──→ j
候选代价 = dist[i][k] + dist[k][j]
```

| 对比 | Bellman–Ford 快照版本 | Floyd–Warshall |
| --- | --- | --- |
| 通常求什么 | 一个起点到所有节点 | 所有节点对之间 |
| DP 状态 | 最多 t 条边到达 v 的最短距离 | 只允许指定集合中的中间节点时，i 到 j 的最短距离 |
| 外层新增的允许范围 | 多允许一条边 | 多允许一个中间节点 k |
| 保留原结果的含义 | 不增加边 | 不经过新中间节点 k |
| 拼接方式 | 已知路径＋一条边 | 两段已知路径 |

令 `D_t[v]` 表示最多 t 条边的最短距离：

```text
D_t[v] = min(D_(t-1)[v], 所有入边 u→v 的 D_(t-1)[u] + w(u,v))
```

令 `F_k[i][j]` 表示只允许编号 `0..k-1` 的节点作为中间节点时的最短距离，`F_0` 由空路径和直接边初始化。允许节点 k 后：

```text
F_(k+1)[i][j] = min(F_k[i][j], F_k[i][k] + F_k[k][j])
```

**Floyd 中的 k 是新增中间节点的编号，不是路径允许使用的边数。** 这里限制的是内部节点，端点 i、j 不受这个编号范围限制。上述有限最短距离的解释以不存在相关负环为前提。

<a id="relax-via"></a>

### 6.1 可复用的 RelaxVia

共同提取“判断可达 → 相加 → 严格改善时更新”即可，不必强行合并两种外层调度。

```go
// RelaxVia 拼接两段路径，返回是否改善了当前答案。
// 约定：inf 大于所有合法有限距离；两段有限代价相加不会溢出。
func RelaxVia(current *int64, left, right, inf int64) bool {
    if left == inf || right == inf {
        return false
    }
    candidate := left + right
    if candidate >= *current {
        return false
    }
    *current = candidate
    return true
}
```

本函数检查不可达，不负责检查整数溢出。使用时需根据输入约束选择 int64 和足够大的 INF，并保证所有实际执行的加法在范围内；“使用 int64”本身不等于不会溢出。直接边的权重也不能与 INF 混用。

Bellman–Ford 传入“一段距离＋一条边权”：

```go
if RelaxVia(&next[v], dist[u], weight, INF) {
    changed = true
}
```

Floyd–Warshall 传入“两段距离”：

```go
RelaxVia(&dist[i][j], dist[i][k], dist[k][j], INF)
```

<a id="ford-floyd-loops"></a>

### 6.2 不同的外层循环

以下片段依赖上述 RelaxVia、已初始化的 int64 距离表，以及包含 from、to、weight 字段的边；它们是外层骨架，不是单独的可执行文件。

#### Bellman–Ford：按边数逐轮传播

```go
for round := 0; round < vertexCount-1; round++ {
    next := append([]int64(nil), dist...)
    changed := false
    for _, edge := range edges {
        if RelaxVia(&next[edge.to], dist[edge.from], edge.weight, INF) {
            changed = true
        }
    }
    dist = next
    if !changed {
        break
    }
}
```

这是“最多使用若干条边”的快照版本，一轮只读旧 dist，写入 next。复制旧表意味着保留不增加边的选择。若改为新表全部初始化 INF，则表示“恰好多用一条边”，此时不能直接套用相同的无改善提前结束逻辑。

普通 Bellman–Ford 也允许原地松弛，但原地一轮可能连续传播多条边，不能严格解释成只增加一条边的状态层。限制总边数时必须使用快照，不能依赖边的扫描顺序。

#### Floyd–Warshall：按中间节点扩展

```go
for k := 0; k < vertexCount; k++ {
    for i := 0; i < vertexCount; i++ {
        for j := 0; j < vertexCount; j++ {
            RelaxVia(&dist[i][j], dist[i][k], dist[k][j], INF)
        }
    }
}
```

**标准 Floyd DP 必须把 k 放在最外层。** 不能随意改成 i、j、k，只做一次三重循环就认为结果相同。

无负环时 `dist[k][k]=0`，本轮经过 k 不会改善拼接基础 `dist[i][k]`、`dist[k][j]`，所以二维表可以原地更新，不必复制整张上一层表。这与限边 Bellman–Ford 必须隔离新旧状态的原因不同。

也不能因为某个 k 的整轮没有改善，就照搬 Bellman–Ford 的提前结束：后面尚未允许的中间节点仍可能产生更短路径。例如边 `0→2=1、2→1=1、0→1=10`，k=0 没有改善，但 k=2 会把 0 到 1 的距离改善为 2。

<a id="ford-floyd-summary"></a>

### 6.3 初始化、负环与复杂度

共同的初始化原则是：**空路径为 0，未知路径为 INF，直接边取最小值。**

- Bellman–Ford：一维表只设 `dist[start]=0`，其余为 INF；边通过后续扫描处理，无需先将全部直达边写入表。
- Floyd–Warshall：先填 INF，再设 `dist[i][i]=0`，最后写入每条直接边 `dist[u][v]=min(dist[u][v], weight)`。这个顺序保留负自环，同时处理重边；无向图要写入两个方向。

两者都允许负权边；如果存在影响某对端点的负环，那对端点之间不存在有限的最短距离。

| 对比 | Bellman–Ford | Floyd–Warshall |
| --- | --- | --- |
| 负环检测 | V-1 轮后仍有起点可达的边可以改善 | 完成所有 k 后，存在 dist[k][k]<0 |
| 检测范围 | 起点可达的负环 | 图中任意位置的负环 |
| 受负环影响的范围 | 起点可到达负环，且负环能继续到达的节点 | 存在负环节点 k，i 可达 k 且 k 可达 j 的节点对 |
| 时间 | 原地版 O(VE)；本节快照版还需 O(V²) 的逐轮复制，即 O(VE+V²) 上界 | O(V³) |
| 距离辅助空间 | O(V)，快照版两张一维表 | O(V²) |

无论使用哪一种算法，“发现负环”和“标记所有受影响的最短距离”为两项工作。仅检查是否存在负环，不会自动把相关结果变成负无穷；不能把负环影响下的表中数值当作最终最短距离。

**可复用的是局部松弛内核：两段可达路径相加，再取 min。应保留的是各自外层不变量：Bellman–Ford 逐步扩展边数预算，Floyd–Warshall 逐步扩展允许的中间节点集合。**
