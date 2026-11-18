# 图算法：先确定问题，再选择模板

[返回仓库目录](../README.md)

同样使用队列或递归，正确性依据可能完全不同。先写出**状态、边的含义、目标、不变量**，再选择下表中的骨架。

| 要解决的问题 | 模板与目录 | 核心不变量 | 本仓库题目 |
| --- | --- | --- | --- |
| 等权最少步数、同步扩散 | [BFS / 多源 / 状态 BFS](bfs/README.md) | 首次入队即确定完整状态的最短距离 | [1091](https://leetcode.com/problems/shortest-path-in-binary-matrix/)、[994](https://leetcode.com/problems/rotting-oranges/)、[542](https://leetcode.com/problems/01-matrix/) 等；24 题内联相同 BFS 模板 |
| 连通块、边界可达区域 | [网格 FloodFill](dfs/grid/readme.md) | 标记后不撤销；每格至多处理一次 | [733](https://leetcode.com/problems/flood-fill/)、[200](https://leetcode.com/problems/number-of-islands/)、[695](https://leetcode.com/problems/max-area-of-island/)、[1020](https://leetcode.com/problems/number-of-enclaves/)、[1254](https://leetcode.com/problems/number-of-closed-islands/)、[130](https://leetcode.com/problems/surrounded-regions/)、[1905](https://leetcode.com/problems/count-sub-islands/)、[827](https://leetcode.com/problems/making-a-large-island/) |
| 一般可达、树上汇总、路径枚举 | [DFS 模板族](dfs/README.md) | 全局访问、子树返回值、当前路径三种状态分别维护 | [1971](https://leetcode.com/problems/find-if-path-exists-in-graph/)、[2316](https://leetcode.com/problems/count-unreachable-pairs-of-nodes-in-an-undirected-graph/)、[1443](https://leetcode.com/problems/minimum-time-to-collect-all-apples-in-a-tree/)、[1519](https://leetcode.com/problems/number-of-nodes-in-the-sub-tree-with-the-same-label/)、[79](https://leetcode.com/problems/word-search/)、[131](https://leetcode.com/problems/palindrome-partitioning/)、[95](https://leetcode.com/problems/unique-binary-search-trees-ii/) 等 |
| 无向连通性、等价类合并 | [并查集](union_find/README.md) | 集合代表相同即属于同一分量 | [547](https://leetcode.com/problems/number-of-provinces/)、[684](https://leetcode.com/problems/redundant-connection/)、[1319](https://leetcode.com/problems/number-of-operations-to-make-network-connected/)、[721](https://leetcode.com/problems/accounts-merge/)，各题含本地完整集合模板 |
| 前置依赖、DAG 状态传播 | [拓扑排序](topo_sort/readme.md) | 全部前置贡献就绪后才能出队 | [207](https://leetcode.com/problems/course-schedule/)、[210](https://leetcode.com/problems/course-schedule-ii/)、[2115](https://leetcode.com/problems/find-all-possible-recipes-from-given-supplies/)、[1462](https://leetcode.com/problems/course-schedule-iv/)、[2050](https://leetcode.com/problems/parallel-courses-iii/)、[1857](https://leetcode.com/problems/largest-color-value-in-a-directed-graph/)、[329](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/)、[2192](https://leetcode.com/problems/all-ancestors-of-a-node-in-a-directed-acyclic-graph/) |
| 非负权单源最优路径 | [Dijkstra](single_source_shortest_path/README.md) | 弹出的非过期最优标签可以确定 | [743](https://leetcode.com/problems/network-delay-time/)、[1631](https://leetcode.com/problems/path-with-minimum-effort/)、[1514](https://leetcode.com/problems/path-with-maximum-probability/)，按标签定义适配本地 BestLabels |
| 负边、边数限制 | [Bellman-Ford 与边数分层](single_source_shortest_path/README.md) | 第 k 轮只读取上一轮的有限边数状态 | 通用示例、[787](https://leetcode.com/problems/cheapest-flights-within-k-stops/) |
| 所有点对路径 | [Floyd-Warshall](all_pairs_shortest_path/Readme.md) | 第 k 阶段只允许前 k 个点作为中间点 | 通用距离、[399](https://leetcode.com/problems/evaluate-division/) 比值实现；[1334](https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/) 保留片段 |
| 无向图最小连通总代价 | [Kruskal / Prim](minimum_spanning_tree/readme.md) | 安全地把不同分量连接起来 | 通用示例、[1584](https://leetcode.com/problems/min-cost-to-connect-all-points/) |
| 相邻节点分两组 | [二分图染色](bipartite/README.md) | 每条边两端颜色相反 | [785](https://leetcode.com/problems/is-graph-bipartite/) |
| 删除边后是否断开 | [桥与割点](bridge_articulation_point/README.md) | 子树能否绕过父边回到祖先 | [1192](https://leetcode.com/problems/critical-connections-in-a-network/) |
| 从每个起点分别求有向覆盖 | [有向连通与覆盖](connectivity/README.md) | 每次搜索有独立访问表 | [2101](https://leetcode.com/problems/detonate-the-maximum-bombs/) |

## 同一模板的本地副本与参考实现

- [FloodFill](dfs/grid/flood_fill.go)：八道网格题各自内联完整四方向遍历，题目只决定可进入条件和访问后的汇总操作；本链接保留参考实现及测试。
- [Kahn](topo_sort/template/kahn.go)：[207](https://leetcode.com/problems/course-schedule/) 与六题主线各自内联入度、队列与回调骨架，题目负责建图、种子与 DP 转移。配方题的种子来自供应品，不能机械选择全部零入度名字。
- BFS、并查集、BestLabels 和 FloydClosure 的 README 也只声明一次模板，再展示适配函数；实际每题主文件包含其使用的完整辅助声明，没有仓库内部导入。

其他算法保持题目内的独立状态与类型。例如单词搜索必须恢复路径访问标记，岛屿统计不能恢复；Prim 的接入边权也不等于 Dijkstra 的源点距离。

有现有测试的网格八题、拓扑主线及 [207](https://leetcode.com/problems/course-schedule/)、并查集四题、三种 Dijkstra 标签均已按模板实现；BFS 其余扩展题和通用示例的验证范围见各族说明。编译成功不等于每个无测试历史片段都已穷尽验证。

## 练习与原始资料

优先练习路径：[200](https://leetcode.com/problems/number-of-islands/) → [1091](https://leetcode.com/problems/shortest-path-in-binary-matrix/) → [994](https://leetcode.com/problems/rotting-oranges/) → [785](https://leetcode.com/problems/is-graph-bipartite/) → [547](https://leetcode.com/problems/number-of-provinces/) → [207](https://leetcode.com/problems/course-schedule/) → [210](https://leetcode.com/problems/course-schedule-ii/) → [743](https://leetcode.com/problems/network-delay-time/) → [1631](https://leetcode.com/problems/path-with-minimum-effort/) → [2050](https://leetcode.com/problems/parallel-courses-iii/) → [827](https://leetcode.com/problems/making-a-large-island/)。

目录中的 `variants/` 保存同题其他写法；无扩展名文件仍是原始笔记，不会被 Go 自动编译。明确题意的 `.go` 练习按本地模板实现，历史材料的状态与保留边界在对应 README 说明。[迁移清单](MIGRATION.json)逐文件记录旧路径和新路径；[原图算法笔记](NOTES.md)保留原始外部资料链接。
