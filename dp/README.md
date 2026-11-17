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

这些是共同的推导步骤，不是一条可以不加判断套用的转移公式。各家族 README 先给出可执行子模板，再展示每道题的实际适配函数。题目源文件内嵌相同辅助函数，可以单文件复制到 LeetCode；选择一种适配器时，也复制其列出的模板一次。

## 归并结果与交叉题型

- 原 triangle 的网格题进入 grid；爬楼梯进入 linear。
- 原 Interval 与 merge_interval 进入 interval；石子博弈及区间分差变体集中到 game。
- 原 distinct_ways 按实际状态分别进入 grid、knapsack、linear。
- 原 coin 的零钱题进入 knapsack；原 house_robber 的线性两题进入 linear。
- 树形打家劫舍的独立 DP 版本位于 [tree/postorder/l337/variants/dp](../tree/postorder/l337/variants/dp/337.go)，保留原测试。
- 原 knapsack/871 实际是大顶堆贪心，保存在 [heap_priority_queue/l871/legacy/from_knapsack/871](../heap_priority_queue/l871/legacy/from_knapsack/871)。
- grid 中的 391 是历史几何草稿，单独标记为 [legacy_geometry](grid/legacy_geometry/l391/391.go)，不属于网格 DP 模板。
- 字符串回文区间题保留原实现位置，由 interval 交叉引用。双序列题目也在各题文件内嵌所需模板，公共 template.go 保留用于单独学习与验证。

逐文件迁移记录见 [MIGRATION.json](MIGRATION.json)。[原 DP 总笔记](legacy/README.original.md)及各家族 legacy 中的原笔记继续保留。

## 当前实现与历史材料

本轮除整理说明外，也把题目源文件改写为对应子模板的适配器，补全了 300、343、1463、1340 和原 stone 的 1690 练习，修正了 714 的手续费状态转移。原有测试保留，新增边界测试覆盖原先没有测试的已补全实现。

variants 继续保留不同方法，例如 741／1463 的分层与递归、486 的分差与双方收益、1690 的递归与区间迭代。legacy 中的历史解释和无扩展名片段仍可查阅；各家族改写前的完整 README 保存在 `legacy/README.before-template-rewrite.md`，其中旧的实现状态描述不代表当前源文件。

几何题 391 独立保留在 legacy_geometry；它的面积与角点奇偶性不属于网格 DP，不会套用本系列递推。

```sh
go test ./dp/...
go test ./dp/grid/l64 ./dp/knapsack/l416 ./dp/interval/l312
```
