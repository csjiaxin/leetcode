# Dijkstra：非负边权最短路径

[返回模板目录](../README.md)

## 目录

- [条件与不变量](#invariant)
- [Go 通用骨架](#code)
- [松弛变化点与例子](#mapping)
- [复杂度与陷阱](#pitfalls)

<a id="invariant"></a>

## 1. 条件与不变量

标准版本要求所有边权非负。维护源点到每个节点的最佳已知距离，最小堆优先处理距离最小者。与 BFS 相比，不能因第一次入堆就标记为最终确定。

**不变量**：每次弹出一个非过期的最小距离，它就是该节点的最短距离。非负边保证尚未处理的更远节点无法绕路产生更短结果。

同一节点可能先以较大距离入堆，后来又以更小距离入堆；不必删除旧堆项，弹出时比较 `item.distance!=dist[item.node]`，跳过过期项。

<a id="code"></a>

## 2. Go 通用骨架

```go
package template

import "container/heap"

const Infinity int64 = 1<<63-1

type Edge struct { To int; Weight int64 }
type Item struct { Node int; Distance int64 }
type MinHeap []Item
func(h MinHeap) Len() int { return len(h) }
func(h MinHeap) Less(i,j int) bool { return h[i].Distance<h[j].Distance }
func(h MinHeap) Swap(i,j int) { h[i],h[j]=h[j],h[i] }
func(h *MinHeap) Push(x any) { *h=append(*h,x.(Item)) }
func(h *MinHeap) Pop() any {
    old:=*h
    x:=old[len(old)-1]
    *h=old[:len(old)-1]
    return x
}

// 节点编号合法，src 存在，所有边权非负。
// 所需有限距离必须 < Infinity；更大的结果不在本模板表示范围内。
func Dijkstra(graph [][]Edge,src int) []int64 {
    dist:=make([]int64,len(graph))
    for i:=range dist { dist[i]=Infinity }
    dist[src]=0
    queue:=&MinHeap{{Node:src,Distance:0}}
    heap.Init(queue)
    for queue.Len()>0 {
        current:=heap.Pop(queue).(Item)
        u,d:=current.Node,current.Distance
        if d!=dist[u] { continue }
        // 若只求某个 target，在过期检查后可于 u==target 时返回。
        for _,e:=range graph[u] {
            if e.Weight<0 { panic("negative edge weight") }
            if e.Weight>Infinity-d { continue } // 防止加法溢出
            candidate:=d+e.Weight
            if candidate<dist[e.To] {
                dist[e.To]=candidate
                heap.Push(queue,Item{Node:e.To,Distance:candidate})
            }
        }
    }
    return dist
}
```

边权非负是调用者必须保证的前提，代码中的检查只发现实际遍历到的违规边。Infinity 同时是不可达标志，不用它表示真实路径值；根据题目范围选安全类型和哨兵。

<a id="mapping"></a>

## 3. 松弛变化点与例子

```text
边：0→1(10)，0→2(1)，2→1(1)
弹出 (0,0)：1 更新为10，2更新为1
弹出 (2,1)：1 更新为2
弹出 (1,2)：距离确定为2
弹出 (1,10)：10!=dist[1]，跳过
```

固定过程：堆取当前最优 → 跳过过期项 → 遍历邻居 → 仅在严格改善时更新并入堆。

| 题目 | 路径标签和初值 | 候选值 | 堆与比较 | 最终答案 |
| --- | --- | --- | --- | --- |
| [743. 网络延迟时间](https://leetcode.com/problems/network-delay-time/) | 累计距离，源=0、其他=INF | d+w | 最小堆，candidate<dist[v] | 所有距离最大值，有不可达则 -1；节点编号转 0-based |
| [1631. 最小体力消耗路径](https://leetcode.com/problems/path-with-minimum-effort/) | 路径最大相邻高度差，源=0 | max(d,abs(height[u]-height[v])) | 最小堆，candidate<dist[v] | 右下角标签 |
| [1514. 概率最大的路径](https://leetcode.com/problems/path-with-maximum-probability/) | 路径成功概率，源=1、其他=0 | probability×edgeProbability | 最大堆，candidate>best[v] | 终点最大概率，不可达为 0 |

1631、1514 不是直接替换输入边权就套标准加法：必须同步修改标签含义、初值、松弛和堆比较。1514 要求概率在 [0,1]，沿路径扩展不会增加概率，对应最大优先的支配关系；一般自定义代价未必满足这个性质。

<a id="pitfalls"></a>

## 4. 复杂度与陷阱

- 本懒删除堆版本时间 O(V+E log(E+2))，空间 O(V+E)，堆可能存 O(E) 条目。简单图中常写 O((V+E)log(V+1))。
- 最小堆不是按节点编号排序，而是距离。struct 字段不要把 Node 和 Distance 放反。
- 必须跳过旧距离，否则重复扩展大量边；只在严格改善时入堆，避免零权环无限添加等值项。
- 不要“首次入队即 visited”，上面的 0→1 例子会漏掉更短路线。
- 建图要区分有向和无向：743 的边有方向；1514 的边无方向，要添加两边。
- 不能用普通 Dijkstra 处理负边；有向无环图即使有负边也可考虑拓扑 DP，其他场景考虑 Bellman-Ford。
- 不需要为 1514 随意加入 epsilon 阻止微小改善；答案容差与松弛条件是不同概念。

已有练习：[最短路径](../../graph/SingleSourceShortestPath/)、[743](../../heap_priority_queue/l743/)、[1631](../../heap_priority_queue/l1631/)、[1514](../../heap_priority_queue/l1514/)。
