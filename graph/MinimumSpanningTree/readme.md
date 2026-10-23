https://leetcode.com/tag/minimum-spanning-tree/

**最小生成树（Minimum Spanning Tree，MST）**用于无向带权图：选择一些边，连接所有节点，使总权重最小。

算法	核心思路	常见时间复杂度
Kruskal	按边权排序，用并查集合并不同集合	O(E log E)
Prim	从一个节点开始，每次用最便宜的边接入一个新节点	数组版 O(V²)；堆版 O(E log V)


一 Prim

Prim:
key[v] = 将 v 接入当前生成树所需的最小边权
更新：key[v] = min(key[v], weight(u, v))

Dijkstra:
dist[v] = 从起点到 v 的最短路径长度
更新：dist[v] = min(dist[v], dist[u] + weight(u, v))


二 Kruskal
对于 n 个节点，生成树：
- 包含 n-1 条边。
- 连接所有节点。
- 不包含环。
可以使用 Kruskal 算法：边按权重排序 + 并查集判环。
1. 按边权从小到大处理。
2. 如果两个端点属于不同集合，就选择这条边并合并集合。
3. 选满 n-1 条边后结束。