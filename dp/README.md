# 动态规划：按状态和依赖关系归类

本目录按可复用的状态模型归并为八个题型家族。题目有多个标签时，以一处实现为准，用链接交叉索引；不要求所有题共用同一条递推式。

| 家族 | 状态与计算顺序 | 模板入口 |
| --- | --- | --- |
| 线性与有限状态 | 前缀、结尾、选择状态，或位置加额外状态 | [linear](linear/README.md) |
| 网格与分层状态 | 单格、区域、步数或多个位置；依赖顺序由转移决定 | [grid](grid/README.md) |
| 背包 | 已处理物品与容量；区分 0/1、重复使用、组合与排列 | [knapsack](knapsack/README.md) |
| 区间 | 先计算短区间，再计算依赖它的长区间 | [interval](interval/README.md) |
| 博弈 | 当前行动者的必胜性、分差或可得分数 | [game](game/README.md) |
| 股票状态机 | 日期、持股状态、次数及冷冻等限制 | [stock](stock/README.md) |
| 双序列与字符串 | 两个前缀、回文区间、分割及模式匹配 | [strings](strings/readme.md) |
| 跳跃题族对照 | 同题族中的贪心、BFS、DAG DP，分别选模板 | [jump_game](jump_game/README.md) |

## 从题意到模板

先写清六件事：

1. 状态包含什么信息；哪些信息能由其他维度推出。
2. 状态值代表最大值、最小值、方案数还是可行性。
3. 空状态、不可达状态与终止状态如何初始化。
4. 哪些前驱合法，合并时取 max、min、加法还是逻辑或。
5. 遍历顺序是否保证所有依赖已计算；压缩空间后是否会重复使用当前物品。
6. 最终答案来自一个状态、全部状态的汇总，还是还需要恢复实际方案。

这些是共同的推导步骤，不是一条可以不加判断套用的转移公式。可执行公共辅助函数目前保留在 [strings/template.go](strings/template.go)：它只承诺双前缀表格的边界和行优先计算顺序。其他家族先用各自 README 的骨架，避免把网格、区间和博弈的不同不变量强行合并。

## 归并结果与交叉题型

- 原 triangle 的网格题进入 grid；爬楼梯进入 linear。
- 原 Interval 与 merge_interval 进入 interval；石子博弈及区间分差变体集中到 game。
- 原 distinct_ways 按实际状态分别进入 grid、knapsack、linear。
- 原 coin 的零钱题进入 knapsack；原 house_robber 的线性两题进入 linear。
- 树形打家劫舍的独立 DP 版本位于 [tree/postorder/l337/variants/dp](../tree/postorder/l337/variants/dp/337.go)，保留原测试。
- 原 knapsack/871 实际是大顶堆贪心，保存在 [heap_priority_queue/l871/legacy/from_knapsack/871](../heap_priority_queue/l871/legacy/from_knapsack/871)。
- grid 中的 391 是历史几何草稿，单独标记为 [legacy_geometry](grid/legacy_geometry/l391/391.go)，不属于网格 DP 模板。
- 字符串回文区间题保留原实现位置，由 interval 交叉引用。双序列系列和 leetcode/dp/strings 导入路径保持稳定。

逐文件迁移记录见 [MIGRATION.json](MIGRATION.json)。[原 DP 总笔记](legacy/README.original.md)及各家族 legacy 中的原笔记继续保留。

## 历史实现与验证边界

目录归并保留原代码、测试、替代方法、无扩展名片段及已有调试文件。variants 表示不同实现；legacy 表示历史材料，名称不保证实现完整或已通过测试。Go 不会把原有无扩展名片段当作 Go 源文件，这一行为没有改变。

未完成的 .go 草稿和原有会失败的测试仍照常参与构建/测试，没有通过改扩展名、构建标签或隐藏目录排除它们。因此全量命令可能继续报告原有失败：

~~~sh
go test ./dp/...
~~~

练习某一题时可只运行对应包，例如：

~~~sh
go test ./dp/strings/...
go test ./dp/grid/l64 ./dp/knapsack/l416 ./dp/interval/l312
~~~

本轮是目录与模板文档整理，不把历史草稿自动改写成另一套算法，也不把单纯的路径迁移当成正确性修复。

本轮全量 DP 检查仍报告以下已有问题；这些实现与测试均未修改：

| 包 | 当前问题 |
| --- | --- |
| game/l1690/legacy/from_stone | 未实现，运行时 panic |
| grid/l1463 | 未实现，运行时 panic |
| grid/legacy_geometry/l391 | 缺少返回值 |
| linear/l300 | 未完成，存在未使用变量和缺少返回值 |
| linear/l343 | 未完成，引用未定义标识符 |
| jump_game/l1340 | 原草稿返回 0，原有样例与边界测试失败 |
| stock/l714 | 原实现未正确处理手续费，原有测试失败 |
