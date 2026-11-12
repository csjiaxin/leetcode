# 算法通用模板：固定骨架＋变化点

目标不是给所有题硬套同一段代码，而是掌握**共同思路，再看每道题如何落地**。本目录共10个专题、37组练习；每道题的说明和实现分开，不把多题实现塞进同一个代码块。代码只放在文档中，不填充或覆盖练习函数。

## 阅读方式

每份专题README采用相同结构：

1. **核心思路**：适用条件、不变量和文字骨架，不额外封装复杂回调。
2. **题目导航**：目录直接跳到对应题目。
3. **逐题独立章节**：原题链接 → 本地练习/测试 → 题意 → 状态与推导 → 完整Go代码 → 示例与边界 → 复杂度。
4. **对比与易错点**：总结不同题目究竟改了什么。

每个 `go` 代码块只实现一道题，包含该题需要的import和辅助类型，可以单独复制。堆和并查集辅助代码可能在不同题中重复，这是为避免阅读时跨章节找依赖。不要把多个完整代码块直接拼接到同一个Go文件。

可先自己完成练习，再展开或对照文档中的实现。现有练习可能已由你填写；本次整理只改README，不重置任何实现。

## 目录与学习路线

| 顺序 | 模板 / 专题 | 何时考虑 | 固定骨架 | 主要变化点 |
| --- | --- | --- | --- | --- |
| 1 | [前缀状态＋哈希表](prefix_state_hash/README.md) | 连续区间的和、余数、数量差、奇偶性 | 更新状态 → 查询历史 → 记录当前 | 状态定义、查询键、次数或最早下标 |
| 2 | [不定长滑动窗口](sliding_window/README.md) | 连续区间，移动左端点能单调修复限制 | 右端加入 → 收缩 → 更新答案 | 加入/移除状态、合法条件、求长/求短 |
| 3 | [atMost 相减计数](at_most/README.md) | 恰好 K 类 / K 次事件的子数组 | atMost(K)-atMost(K-1) | 事件定义、窗口状态、负预算处理 |
| 4 | [单调栈](monotonic_stack/README.md) | 最近更大/更小位置、贡献边界 | 遍历 → 弹栈结算 → 入栈 | 比较符号、相等归属、结算公式 |
| 5 | [BFS / 多源 BFS](bfs/README.md) | 每步代价相同的最短路或传播 | 初始化源 → 入队标记 → 扩展 | 状态、邻居、源点、距离单位 |
| 6 | [Dijkstra](dijkstra/README.md) | 非负边权最短路 | 堆取最优 → 跳过过期 → 松弛 | 建图、路径代价定义、堆方向 |
| 7 | [并查集 DSU](union_find/README.md) | 合并集合、无向连通性 | Find＋Union | 节点映射、分量信息、失败合并的含义 |
| 8 | [固定大小 K 的堆](top_k_heap/README.md) | 数据流前 K 大 / 小 | 保留 K 个候选 → 淘汰最差 | 比较键、频率、是否需要排序输出 |
| 9 | [拓扑排序](topological_sort/README.md) | 有向依赖、先后顺序、判环 | 零入度入队 → 删出边 | 建边方向、返回顺序或 DAG DP |
| 10 | [差分＋前缀扫描](difference_array/README.md) | 批量区间加减、覆盖和同时占用 | 两端记变化 → 前缀还原 | 闭/半开端点、坐标范围、聚合目标 |

## 推荐练习链（链接直达 leetcode.com）

- 前缀状态：[560](https://leetcode.com/problems/subarray-sum-equals-k/) → [974](https://leetcode.com/problems/subarray-sums-divisible-by-k/) → [525](https://leetcode.com/problems/contiguous-array/) → [1371](https://leetcode.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/)。优先练这一组，体会“次数”和“最早位置”的区别。
- 滑窗：[3](https://leetcode.com/problems/longest-substring-without-repeating-characters/) → [209](https://leetcode.com/problems/minimum-size-subarray-sum/) → [76](https://leetcode.com/problems/minimum-window-substring/)。求最长与求最短的答案更新时机不同。
- 恰好 K：[992](https://leetcode.com/problems/subarrays-with-k-different-integers/) → [1248](https://leetcode.com/problems/count-number-of-nice-subarrays/)。也可与前缀状态计数对照。
- 栈：[739](https://leetcode.com/problems/daily-temperatures/) → [84](https://leetcode.com/problems/largest-rectangle-in-histogram/) → [907](https://leetcode.com/problems/sum-of-subarray-minimums/)。最后一题重点是重复值归属。
- 图：[1091](https://leetcode.com/problems/shortest-path-in-binary-matrix/) → [994](https://leetcode.com/problems/rotting-oranges/) → [743](https://leetcode.com/problems/network-delay-time/)。先确认边的代价，再选择队列还是堆。
- 集合与堆：[547](https://leetcode.com/problems/number-of-provinces/)、[703](https://leetcode.com/problems/kth-largest-element-in-a-stream/)；依赖与扫描：[210](https://leetcode.com/problems/course-schedule-ii/)、[1094](https://leetcode.com/problems/car-pooling/)。

## 如何识别模板，而不是背题号

1. 写出状态与合法性：子数组？路径？集合？区间事件？
2. 证明关键性质：窗口能否单调收缩、check 是否单调、路径扩展会不会变优？
3. 对照不变量填骨架，只改题目需要的部分。
4. 检查空输入、重复值、负数、不可达、相等边界、数值溢出。
5. 用很小的暴力算法对照测试，验证迁移是否正确。

反例尤其重要：含负数的和限制不一定能滑窗；负权图不能套普通 Dijkstra；DSU 不能处理一般有向判环；Top K 堆不能独自解决中位数；最优化问题不一定存在贪心策略。

## 与已有笔记的关系

- [闭区间 FirstTrue 与答案二分](../binary_search/readme.md)：false→true，题目提供 check。
- [贪心系列](../greedy/readme.md)：必须证明局部选择不会损害全局解。
- [动态规划](../dp/README.md)：按背包、区间、树形等分类，状态和转移不能完全黑盒化。
- 各专题会链接原仓库的练习目录，不搬动旧实现，不复制题目答案进空函数。

## 使用代码骨架

各题Go代码块以 `package template` 开头，函数签名与对应练习一致；可单独复制为 `.go` 文件。它们是**文档参考实现**，不写入练习源码。使用Go 1.21+（内置min/max、any）；计数、前缀和、距离按题目范围使用int64中间值。复杂度中的哈希操作默认为平均O(1)。

## 练习目录与测试

每个专题 README 中链接的题目都有独立的 `l题号/题号.go` 和 `l题号/题号_test.go`，共 37 组。跨专题出现的同题保留独立练习版本，例如 1248；295 是 Top K 专题中的双堆对照题，不属于固定大小 K 堆的直接应用。

**练习文件初始为空函数，文档则提供独立参考实现。** 尚未填写的返回值函数会报 `missing return`，不能据此把整个目录测试失败当作算法错误。完成某题后，单独运行：

```sh
go test ./algorithm_templates/prefix_state_hash/l560 -v
```

测试包含示例和关键边界。拓扑顺序、账户合并、Top K 集合及最近点不要求固定输出顺序；有状态题按操作序列验证。文档中的参考模板和练习文件相互独立，现有仓库题目实现不改动。
