# 拓扑排序：先后依赖与有向判环

[返回模板目录](../README.md)

## 目录

- [条件与不变量](#invariant)
- [Go 固定骨架](#code)
- [变化点与推演](#mapping)
- [复杂度与陷阱](#pitfalls)

<a id="invariant"></a>

## 1. 条件与不变量

有向边 u→v 表示 u 必须先于 v。若有环，不存在满足全部依赖的顺序；有向无环图 DAG 存在拓扑序，可能不唯一。

Kahn 算法维护当前剩余图的入度。**不变量**：队列中节点的前置依赖全部完成；处理 u 后删除其出边，某个 v 的剩余入度降至 0 时才允许入队。

队列只是保存“现在可以做”的任务，不代表每层耗时相同。2050 的不同课程耗时要另外维护最长前驱完成时间，不能只数 BFS 层数。

<a id="code"></a>

## 2. Go 固定骨架

```go
package template

// edges 中 [u,v] 表示 u 必须先于 v；节点编号 [0,n)。
// 返回 (order,true) 表示无环；有环时返回 (nil,false)。
func TopologicalOrder(n int,edges [][2]int) ([]int,bool) {
    graph:=make([][]int,n)
    indegree:=make([]int,n)
    for _,e:=range edges {
        u,v:=e[0],e[1]
        graph[u]=append(graph[u],v)
        indegree[v]++
    }
    queue:=[]int{}
    for u:=0;u<n;u++ {
        if indegree[u]==0 { queue=append(queue,u) }
    }
    order:=make([]int,0,n)
    for head:=0;head<len(queue);head++ {
        u:=queue[head]
        order=append(order,u)
        for _,v:=range graph[u] {
            indegree[v]--
            if indegree[v]==0 { queue=append(queue,v) }
        }
    }
    if len(order)!=n { return nil,false }
    return order,true
}

// LC210 的 [course, prerequisite] 必须转成 prerequisite→course。
func CourseOrder(n int,prerequisites [][]int) []int {
    edges:=make([][2]int,0,len(prerequisites))
    for _,p:=range prerequisites {
        edges=append(edges,[2]int{p[1],p[0]})
    }
    order,ok:=TopologicalOrder(n,edges)
    if !ok { return []int{} }
    return order
}
```

邻接表和入度必须一一对应：若保留重复边，就重复计入度并在处理时重复减；若去重，必须两处同时去重。

<a id="mapping"></a>

## 3. 变化点与推演

课程依赖 [[1,0],[2,0],[3,1],[3,2]] 对应：0→1、0→2、1→3、2→3。

```text
入度 [0,1,1,2]，初始队列 [0]
处理0：1、2入度归零，队列追加1、2
处理1：3入度从2降为1，还不能处理3
处理2：3入度归零，追加3
处理3：结束，得到 [0,1,2,3]
```

[0,2,1,3] 同样合法。不应排序最终 order，那可能破坏依赖；若需要字典序最小拓扑序，用最小堆选择当前零入度节点。

| 题目 | 边方向 | 变化点 | 答案 |
| --- | --- | --- | --- |
| [207. 课程表](https://leetcode.com/problems/course-schedule/) | prerequisite→course | 不需要保存具体顺序，只计处理数 | 处理数=n |
| [210. 课程表 II](https://leetcode.com/problems/course-schedule-ii/) | prerequisite→course | 返回顺序，判环失败为空 | order |
| [2050. 并行课程 III](https://leetcode.com/problems/parallel-courses-iii/) | 输入 prev→next，编号1-based需转换 | 完成时间 DP，finish[v]=max(finish[v],finish[u]+time[v]) | 所有 finish 的最大值 |

2050 初始化每个节点 finish[u]=time[u]。处理每条边更新 finish[v]，并在所有入边处理完、入度降为0后入队。因为课程可无限并行，不是把总耗时求和；还需保证图无环。

固定骨架：建图入度 → 零入度初始化 → 出队 → 遍历出边 → 入度归零入队 → 处理数判环。可在处理出边时添加 DAG DP，但不是替换掉依赖约束。

<a id="pitfalls"></a>

## 4. 复杂度与陷阱

- 基础队列版本时间 O(V+E)，空间 O(V+E)。字典序最小的堆版本额外增加堆操作时间。
- 边方向错误仍可能得到一个序列，却正好违反题意，尤其注意课程输入 [a,b] 表示 b→a。
- 所有初始入度0节点都需入队，孤立节点也必须出现在序列中。
- 有环时不一定初始队列为空；可能先处理无环部分，再卡在环及其后继节点。因此最终必须比较处理数与 n。
- 只在入度降到0时入队，不能每减一次就加入，也不能重复入队。
- DSU 判断两点连通不能替代有向判环；两条方向不同的路径可能共享点但没有有向环。
- 时间较大时 DAG DP 的完成时间使用 int64；课程时长不同时，拓扑层数不是总耗时。

已有练习：[graph/topo_sort](../../graph/topo_sort/)。等权最短步数见 [BFS](../bfs/README.md)。
