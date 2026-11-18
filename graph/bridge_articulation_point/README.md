# 桥与割点：DFS 子树能否绕过父边

[返回图算法目录](../README.md)

适用于无向图。disc[u] 是首次访问时间；low[u] 是 u 的 DFS 子树借助树边和回边能触及的最早访问时间。

**不变量**：孩子 v 完成后，low[v] 已汇总该子树的回边信息。若 low[v] > disc[u]，子树无法绕过树边 u-v 返回祖先，该边是桥。

```go
func dfs(u, parentEdge int) {
    disc[u], low[u] = timer, timer
    timer++
    for _, e := range graph[u] {
        if e.id == parentEdge { continue }
        v := e.to
        if disc[v] == -1 {
            dfs(v, e.id)
            low[u] = min(low[u], low[v])
            if low[v] > disc[u] { /* u-v 是桥 */ }
        } else {
            low[u] = min(low[u], disc[v])
        }
    }
}
```

这是需要配好 disc 初值、timer 和边编号的骨架。按**父边编号**跳过，可以保留平行边形成的回路；用邻接集合会丢失重边。非根割点检查 low[v]>=disc[u]；DFS 根则要单独统计孩子数。

| 题目 | 原始代码 | 状态 |
| --- | --- | --- |
| [1192. 查找集群内的关键连接](https://leetcode.com/problems/critical-connections-in-a-network/) | [1192.go](1192.go) | 原始示例保留，尚未添加测试；对照上面的树边／回边更新规则复盘 |

扫描所有未访问节点以覆盖不连通图。时间 O(V+E)，辅助空间 O(V+E)。

## 原始练习资料（保留）

https://leetcode.com/list/9bm5lhts/

Bridge
[1192](https://leetcode.com/problems/critical-connections-in-a-network/). Critical Connections in a Network
