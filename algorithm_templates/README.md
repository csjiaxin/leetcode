# 算法模板总目录：先看不变量，再选择系列

本仓库按状态、信息流向和合法性条件归并题目。同一题的不同实现集中在题目目录的 `variants` 中。各系列 README 展示固定模板与逐题适配，实际题解内嵌同样的模板；不同不变量保留各自子模板，避免为了统一而改变题意。

## 动态规划的八个家族

详细索引：[DP 总结](../dp/README.md)。

| 系列 | 固定状态 / 依赖骨架 | 每题需要修改的部分 | 入口 |
| --- | --- | --- | --- |
| 线性与有限状态 | 按前缀、结尾位置或有限状态推进 | 相邻选择、LIS 前驱、额外状态维度 | [linear](../dp/linear/README.md) |
| 网格与分层 | 从合法前驱累计当前状态 | 移动方向、障碍、步数、多个位置 | [grid](../dp/grid/README.md) |
| 背包 | 已处理物品＋容量 | 是否重复使用、容量遍历方向、组合或排列 | [knapsack](../dp/knapsack/README.md) |
| 区间 | 由短区间推导长区间 | 分割点、最后操作、合并代价、额外维度 | [interval](../dp/interval/README.md) |
| 博弈 | 以当前行动者为视角求值 | 必胜后继、当前收益、对手优势、合法操作 | [game](../dp/game/README.md) |
| 股票状态机 | 日期＋持仓与交易限制 | 买卖更新顺序、次数、手续费、冷冻 | [stock](../dp/stock/README.md) |
| 字符串与双序列 | 两个前缀或回文区间 | max/min/计数/bool、字符关系、边界 | [strings](../dp/strings/readme.md) |
| 跳跃题族 | 根据状态选择贪心、BFS 或 DAG DP | 可达范围、单位边、是否允许回访 | [jump_game](../dp/jump_game/README.md) |

双前缀 [Table](../dp/strings/readme.md) 统一边界和填表顺序；状态含义、转移和最终答案恢复由每题负责。每个解答文件包含自身需要的模板，不依赖仓库内的公共算法包。

## 遍历、依赖与路径

| 系列 | 核心不变量与模板选择 | 入口 |
| --- | --- | --- |
| BFS | 单位边按层扩展；入队即记录完整状态，区分单源、多源、状态 BFS | [graph/bfs](../graph/bfs/README.md) |
| DFS 与网格连通块 | 进入节点先标记；连通块 visited 永久保留，路径枚举才撤销占用 | [graph/dfs](../graph/dfs/README.md) |
| 拓扑与 DAG DP | 节点入队时所有前驱均已处理，先传播再扣入度 | [topo_sort](../graph/topo_sort/readme.md) |
| 单源路径 | 先定义路径代价和松弛规则，再判断 Dijkstra/Bellman-Ford 的前提 | [single_source_shortest_path](../graph/single_source_shortest_path/README.md) |
| 全源路径 | 逐步扩大允许的中间节点集合，k 为最外层 | [all_pairs_shortest_path](../graph/all_pairs_shortest_path/Readme.md) |
| 最小生成树 | 连通块扩展或跨割最小边；边数与连通性单独检查 | [minimum_spanning_tree](../graph/minimum_spanning_tree/readme.md) |
| 并查集 | 只通过代表元判断连通，合并更新大小或秩 | [union_find](../graph/union_find/README.md) |
| 二分图、桥与割点 | 分别维护颜色约束、发现时间与 low-link | [图算法索引](../graph/README.md) |
| 树的后序与路径 DFS | 子树信息向上返回，或祖先状态向下传递 | [树算法](../tree/readme.md) |

## 序列、区间与候选集合

| 系列 | 统一骨架与不能省略的条件 | 入口 |
| --- | --- | --- |
| 二分边界 | 证明谓词 false→true；返回第一个 true | [binary_search](../binary_search/readme.md) · [可执行 FirstTrue](../binary_search/template.go) |
| 双指针 | 证明移动哪个指针能排除候选；读写、相向、链表分别使用各自骨架 | [two_pointers](../two_pointers/readme.md) |
| 滑动窗口 | 最长合法先修复再统计；最短满足先统计再收缩；固定窗口与 atMost 单独处理 | [sliding_window](../sliding_window/README.md) |
| 前缀状态与差分 | 历史状态频次/最早位置，或端点增量后扫描恢复 | [prefix_sum](../prefix_sum/readme.md) · [可执行 Count/Longest](../prefix_sum/state_hash/template.go) |
| 单调栈 | 每个下标入栈、出栈至多一次；相等元素归属决定比较符 | [monotonic_stack](../monotonic_stack/README.md) |
| 嵌套解析栈 | 未完成的外层上下文保存在栈中，闭合时归约 | [stack](../stack/README.md) |
| 堆 | 候选按优先级取出，结合 Top K、双堆、归并、调度等子模板 | [heap_priority_queue](../heap_priority_queue/readme.md) · [可执行 Queue](../heap_priority_queue/internal/pqueue/queue.go) |
| 贪心 | 明确排序或排除规则，并给出交换论证或支配关系 | [greedy](../greedy/readme.md) |
| 区间贪心 | 按结束点排序，区分端点相接是否冲突 | [interval_greedy](../greedy/interval_greedy/readme.md) |

## 搜索与字符串结构

| 系列 | 共同状态 | 入口 |
| --- | --- | --- |
| 回溯 | 当前部分解＋可选集合；进入选择与撤销必须对称 | [back_tracking](../back_tracking/readme.md) |
| 排列变换 | 后缀单调性、枢轴与排名；枚举排列的问题归回溯 | [permutation](../permutation/readme.md) |
| KMP | 当前匹配的模式前缀长度，失配沿前缀函数回退 | [string_match](../string_match/readme.md) |
| Trie | 当前词典前缀节点；按查询规则选择分支与聚合字段 | [trie](../trie/readme.md) |
| 动态区间聚合 | 单位元、结合运算、叶子更新与覆盖区间合并 | [segment tree](../tree/segment/readme.md) |

相似题可以交叉索引。例如 [1248](https://leetcode.com/problems/count-number-of-nice-subarrays/) 可以用前缀计数或 atMost，[871](https://leetcode.com/problems/minimum-number-of-refueling-stops/) 可以用按次数 DP 或最大堆，[329](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/) 可以用记忆化 DFS 或拓扑 DP。算法不变量不同的版本保留为独立解法，不互相替换测试。

## 面试基础补齐：新增指南

这些入口补充先前分散或尚未成体系的知识，包含分类思路、Go 参考模板和练习题单，不将推荐题单当成已完成的独立实现。

| 系列 | 重点不变量 / 选型 | 入口 |
| --- | --- | --- |
| 哈希 | key 的等价性、查写时序、已处理状态 | [hash_table](../hash_table/README.md) |
| 链表 | 保存后继、连接边界、节点身份 | [linked_list](../linked_list/README.md) |
| 树基础 | 层内批处理、BST 中序、结构编码 | [tree/fundamentals](../tree/fundamentals/README.md) |
| 结构组合 | 将 API 的顺序和复杂度要求拆给不同结构 | [data_structure_design](../data_structure_design/README.md) |
| 排序与选择 | 合并有序区间、三路分区、目标排名 | [sorting](../sorting/README.md) |
| 位运算 | 固定位宽、异或抵消、位集合 | [bit_manipulation](../bit_manipulation/README.md) |
| 数学 | 余数递推、指数二进制分解、数位进位 | [math](../math/README.md) |
| 矩阵 | 坐标映射、剩余边界、防更新污染 | [matrix](../matrix/README.md) |
| 扫描线 | 同坐标事件语义、当前活跃候选 | [sweep_line](../sweep_line/README.md) |
| 树状数组 | lowbit 覆盖块、动态前缀与离散化频次 | [tree/fenwick](../tree/fenwick/README.md) |

## 如何复习一个系列

1. 写清适用条件和状态的准确含义。
2. 写出空输入、空前缀、初始候选或终止局面的含义。
3. 说明循环或递归过程中一直成立的不变量。
4. 先写固定骨架，再只修改合法操作、转移或答案聚合。
5. 在题目的独立包运行测试；不要以目录已整理推断实现已完成。

当前实现与验证说明见 [IMPLEMENTATION.md](IMPLEMENTATION.md)。此前的目录迁移说明见 [REORGANIZATION.md](REORGANIZATION.md)，机器可读路径表见 [reorganization.json](reorganization.json)。
