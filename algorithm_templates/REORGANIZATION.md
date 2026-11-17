# 目录合并、模板边界与验证记录

本页记录此前的目录归并阶段；后续源码内嵌模板、补全实现及 README 代码重组的最新状态见 [IMPLEMENTATION.md](IMPLEMENTATION.md)。下方基线包数和共享实现描述仅对应当时的迁移验证。

## 归并结果

本轮先整理 DP，再将同样的分类方法应用于整个仓库：按状态含义、合法性条件、信息流向和更新顺序归并。不同算法解决同一题时保留独立版本；不因都使用递归、二维数组或者两个下标就强行合并转移。

### DP：14 个旧分类归为 8 个家族

| 原目录 | 新的主系列 | 需要区分的子模板 |
| --- | --- | --- |
| Interval、merge_interval | [interval](../dp/interval/README.md) | 分割点、最后操作、区间外附加状态 |
| triangle、grid，distinct_ways 中的网格题 | [grid](../dp/grid/README.md) | 右下路径、分层转移、双位置、前后缀最值优化 |
| house_robber 的线性题、LongestIncreaseSequence、optimize、部分 distinct_ways | [linear](../dp/linear/README.md) | 相邻选择、LIS、前缀最优、有限状态 |
| knapsack、coin 与相关计数题 | [knapsack](../dp/knapsack/README.md) | 0/1、完全背包、组合数、排列数 |
| stone、game、区间分差题 | [game](../dp/game/README.md) | 胜负、分差、位置与额外限制 M |
| stock | [stock](../dp/stock/README.md) | 持仓、交易次数、冷冻和手续费 |
| strings | [strings](../dp/strings/readme.md) | 双前缀、回文区间、切分、不同星号语义 |
| jump_game | [jump_game](../dp/jump_game/README.md) | 同题族中分别选择贪心、BFS、DAG DP |

337 的树形 DP 变体进入 `tree/postorder/l337/variants/dp`；实际使用堆的历史燃料片段进入堆系列。文件名与实际算法不符的旧稿保留并标注，不把它们错误归为同编号 LeetCode 的完整解答。

### 整仓的相似目录与重复题目

| 旧的分散位置 | 归并后的入口 |
| --- | --- |
| bfs、dfs、union_find，以及原图分类 | [graph](../graph/README.md)：BFS、DFS、并查集与路径算法 |
| 原 bfs/grid 中的递归连通块实现 | [graph/dfs/grid](../graph/dfs/grid/readme.md) |
| Subarray 中的前缀状态题、difference_array | [prefix_sum](../prefix_sum/readme.md)：state_hash 与 difference 两个子系列 |
| two_pointers、permutation 中的窗口题，sliding_window 下重复模板 | [sliding_window](../sliding_window/README.md) 的对应 l题号目录 |
| permutation 中的枚举排列、N 皇后等 | [back_tracking](../back_tracking/readme.md) |
| monotonic_stack 中的括号评分；stack 根部解码文件 | [stack](../stack/readme.md) 的 856、394 |
| monotonic_stack 中的滑动窗口最大值 | `sliding_window/l239`，明确使用带过期淘汰的单调队列 |
| heap_priority_queue、monotonic_stack 的重复 *_template | 同题 `l题号/variants/template`，源码与测试独立保留 |
| greedy 的重复燃料题、分散区间题 | [gas_fuel](../greedy/gas_fuel/readme.md)、[interval_greedy](../greedy/interval_greedy/readme.md) |
| binary_search 根部历史片段与示例 | 对应 l题号的 legacy 或 variants/demo |

树、Trie、KMP、单调栈、嵌套栈、双指针、二分和堆继续保留各自系列，因为它们的核心状态与更新规则不同。每个入口给出可复用骨架和变体，而不是只列出题号。

## 模板的实际复用

| 公共实现 | 本轮处理 | 适用边界 |
| --- | --- | --- |
| [双前缀 Table](../dp/strings/template.go) | 保留七道题的现有公共 API 和测试 | 上、左、左上前驱；其他依赖另选模板 |
| [FloodFill](../graph/dfs/grid/flood_fill.go) | 八份相同遍历抽取为共享实现，原题保留薄封装 | 四连通块；访问标记永久保留 |
| [Kahn](../graph/topo_sort/template/kahn.go) | 六份相同拓扑循环抽取为共享实现 | 前驱全部处理后入队；按题意设置初始种子 |
| [FirstTrue](../binary_search/template.go) | 新增可执行公共边界模板与全部分界位置测试 | 单调 false→true 谓词 |
| [Count / Longest](../prefix_sum/state_hash/template.go) | 新增历史状态聚合模板及五种题型适配测试 | 单键等式匹配；频次与最早位置分别处理 |
| [泛型优先队列](../heap_priority_queue/internal/pqueue/queue.go) | 保留已运行的公共候选容器 | 合法性、松弛和调度规则仍由各题负责 |

其余子系列用各自 README 中的固定骨架和状态对照表学习。对原有尚未完成的练习，模板文档提供推导与修改点，不修改测试、隐藏源码或伪装为已完成解答。

## 保留与迁移核对

- [总迁移表](reorganization.json)记录 266 条仓库相对路径映射，包含文件与目录迁移。
- 分项记录为 [DP](../dp/MIGRATION.json)、[图](../graph/MIGRATION.json)和[序列类](../two_pointers/reorganization.json)。
- 迁移前快照中的 821 个文件在迁移后均可找到；所有原测试文件内容不变。
- 原 Go 文件中只有 14 个公共遍历入口改为委托共享实现，分别是八份 FloodFill、六份 Kahn；各题状态转移和原测试保留。
- 空函数、TODO、旧脚本、不同解法及已有调试材料均保留。不同版本同名符号由独立包隔离，不通过改扩展名、构建标签或重定向测试来绕过原有失败。
- 已有的无关工作区改动不回滚；根模块与嵌套模块的配置不变。

## 验证结果

本轮完整回归中，原先通过的 **186 个包全部保持通过**；加上四个新增模板包，共 **190 个包通过**。全量命令仍有 166 个失败包，均能通过迁移映射追溯到基线中已经失败的包。旧混合示例拆包后包数会变化，因此以原通过包的逐包对应和失败来源为准。

```sh
# 会包含原有未完成练习，不能将其退出码解释为所有题已经完成。
go test -count=1 ./...

# 可独立验证的公共模板。
go test ./binary_search ./prefix_sum/state_hash
go test ./graph/dfs/grid ./graph/topo_sort/template

# 迁移目标、内部 Go 导入及总目录链接。
python3 tools/check_catalog.py
python3 -m unittest discover -s tools -p 'test_*.py'
```

另外已对改动过的公共 Go 模板、十四个委托调用包执行 `go vet`，并检查整个仓库 Markdown 的本地链接及 `git diff --check`。目录检查器默认检查两个总入口，其他文档可通过重复传入 `--markdown 路径` 扩展；它不需要网络或第三方 Python 依赖。

原有失败包括未写返回值、显式 panic 练习、未完成标识符，以及已存在的算法错误。它们的原代码与原测试继续暴露这些问题；本轮的目录整理和模板抽取没有把它们误报为通过。
