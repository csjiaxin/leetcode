# BFS / 多源 BFS：等权最短路

[返回模板目录](../README.md)

## 目录

- [适用条件与不变量](#invariant)
- [Go 固定骨架](#code)
- [题目变化点与推演](#mapping)
- [复杂度与易错点](#pitfalls)

<a id="invariant"></a>

## 1. 适用条件与不变量

图中每一步代价相同，求最少步数、最近源距离或同步传播时间。边权不同时普通 FIFO BFS 不保证最短路：0/1 权可用 0-1 BFS，非负一般权用 Dijkstra，负权需其他算法。

**不变量**：首次发现节点时，已经找到最短距离；队列按距离非递减处理。发现后立即设置距离并入队，不能等出队才标记。

多源 BFS 不是对每个源分别 BFS，而是把所有源同时以距离 0 入队，相当于它们处于同一层。得到到最近源的距离。

状态不一定是一个坐标：钥匙集合、剩余消除次数等也可能属于状态；不同状态不能仅因坐标相同而合并。

<a id="code"></a>

## 2. Go 固定骨架

```go
package template

// 节点编号 [0,len(graph))；不可达返回 -1。
func MultiSourceBFS(graph [][]int, sources []int) []int {
    dist:=make([]int,len(graph))
    for i:=range dist { dist[i]=-1 }
    queue:=[]int{}
    for _,s:=range sources {
        if dist[s]!=-1 { continue } // 去重源点
        dist[s]=0
        queue=append(queue,s)
    }
    for head:=0;head<len(queue);head++ {
        u:=queue[head]
        for _,v:=range graph[u] {
            if dist[v]!=-1 { continue }
            dist[v]=dist[u]+1 // 入队前标记
            queue=append(queue,v)
        }
    }
    return dist
}

// 网格多源骨架，不改变原网格。passable 及邻居规则由题意提供。
func GridBFS(rows,cols int, sources [][2]int, directions [][2]int,
    passable func(int,int) bool,
) [][]int {
    dist:=make([][]int,rows)
    for r:=range dist {
        dist[r]=make([]int,cols)
        for c:=range dist[r] { dist[r][c]=-1 }
    }
    queue:=[][2]int{}
    for _,s:=range sources {
        r,c:=s[0],s[1]
        if r<0 || r>=rows || c<0 || c>=cols || !passable(r,c) || dist[r][c]!=-1 { continue }
        dist[r][c]=0
        queue=append(queue,s)
    }
    for head:=0;head<len(queue);head++ {
        r,c:=queue[head][0],queue[head][1]
        for _,d:=range directions {
            nr,nc:=r+d[0],c+d[1]
            if nr<0 || nr>=rows || nc<0 || nc>=cols { continue }
            if !passable(nr,nc) || dist[nr][nc]!=-1 { continue }
            dist[nr][nc]=dist[r][c]+1
            queue=append(queue,[2]int{nr,nc})
        }
    }
    return dist
}
```

固定过程：初始化 -1 → 所有源入队置 0 → 队头扩展 → 首次访问置距离。变化点：sources、邻居 directions、passable 及最终答案汇总。

<a id="mapping"></a>

## 3. 题目变化点与推演

| 题目 | 源点 | 邻居 / 可通行条件 | 输出与距离单位 |
| --- | --- | --- | --- |
| [1091. 二进制矩阵中的最短路径](https://leetcode.com/problems/shortest-path-in-binary-matrix/) | 左上角 | 八方向，值为 0 | dist[终点]+1，因为路径长度计算格子数；不可达为 -1 |
| [994. 腐烂的橘子](https://leetcode.com/problems/rotting-oranges/) | 所有值为 2 的格子 | 四方向，非零格子 | 所有新鲜橘子距离的最大值；有新鲜橘子不可达则 -1；无新鲜为 0 |
| [542. 01 矩阵](https://leetcode.com/problems/01-matrix/) | 所有值为 0 的格子 | 四方向，所有格子可通行 | 返回 dist 矩阵，每次跨一条边计 1 |

994 中把已有腐烂格子也视为可通行是安全的，它们都在初始源集合并已标记，不会重复加入。使用不可变网格＋dist 方案，无需靠修改 grid 判断传播完成。

四方向为上、下、左、右；八方向再加四个对角方向。1091 可对角移动，不能照搬 994 的 directions。

推演：路径 0—1—2—3—4，以 0、4 为源：

```text
初始：queue=[0,4]，dist=[0,-1,-1,-1,0]
弹出0：发现1，dist[1]=1
弹出4：发现3，dist[3]=1
弹出1：发现2，dist[2]=2
弹出3：2已被发现，不再入队
最终 dist=[0,1,2,1,0]
```

1091 单格 [0]：源到自身边数为 0，但题目路径格子数为 1。994 无新鲜橘子：传播时间为 0，不应多加一分钟。

<a id="pitfalls"></a>

## 4. 复杂度与易错点

- 邻接表图时间 O(V+E)，除输入外空间 O(V)。常数邻居的 R×C 网格时间、额外空间均 O(RC)。
- 入队立即标记，确保每个状态最多入队一次；不能从不同父节点反复加入。
- 若按层统计时间，处理一层前固定队列层尾，不能把新入队元素算进当前层。存 dist 时无需另写层循环。
- 源是障碍、终点是障碍、无源、有不可达状态都要按题意处理。
- 所有源同时入队，不能从第一个源完成搜索后把其余源当成访问过而略过。
- BFS 找的是状态图上的最短步数；状态是否包含额外资源，决定 visited 的维度。

已有相关练习：[图论](../../graph/)、[DFS 网格题](../../dfs/)。带权问题见 [Dijkstra](../dijkstra/README.md)。
