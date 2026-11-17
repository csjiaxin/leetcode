# leetcode

## 按模板查找题目

[算法模板总目录](algorithm_templates/README.md)汇总各系列的适用条件、不变量、固定骨架和修改点；[目录合并记录](algorithm_templates/REORGANIZATION.md)说明旧目录的去向。

| 入口 | 覆盖的子系列 |
| --- | --- |
| [动态规划](dp/README.md) | 线性、网格、背包、区间、博弈、股票、字符串与跳跃题族 |
| [图算法](graph/README.md) | BFS、DFS、拓扑、路径、生成树、并查集、二分图与连通性 |
| [树](tree/readme.md) | 后序状态返回、路径状态传递、线段树、平衡搜索树 |
| [序列操作](two_pointers/readme.md) | 双指针；另见[滑动窗口](sliding_window/README.md)、[前缀与差分](prefix_sum/readme.md) |
| [候选维护](heap_priority_queue/readme.md) | 堆、Top K、归并、调度；另见[单调栈](monotonic_stack/README.md)和[嵌套解析栈](stack/README.md) |
| [边界与选择](binary_search/readme.md) | 二分边界；另见[贪心](greedy/readme.md) |
| [枚举与变换](back_tracking/readme.md) | 回溯；另见[排列变换](permutation/readme.md) |
| [字符串结构](string_match/readme.md) | KMP；另见[Trie](trie/readme.md) |

同题的不同实现集中在 `l题号/variants`，原练习、历史片段和测试均保留。未完成函数依然参与测试；目录整理不等于把整个题库的所有练习补全。

在仓库根目录使用 Go 模块模式。例如：

```sh
go test ./dp/strings/...
go test ./binary_search ./prefix_sum/state_hash
python3 tools/check_catalog.py
```

全量 `go test ./...` 包含原有空函数和历史草稿，已知会有失败；本次整理的验证边界见迁移记录。

## 原有参考资料

leetcode support github.com/emirpasic/gods/stacks/arraystack
Support https://pkg.go.dev/github.com/emirpasic/gods@v1.18.1 and https://pkg.go.dev/github.com/emirpasic/gods/v2@v2.0.0-alpha library.


Monotonic Stack(单调栈)
https://github.com/labuladong/fucking-algorithm/blob/master/%E6%95%B0%E6%8D%AE%E7%BB%93%E6%9E%84%E7%B3%BB%E5%88%97/%E5%8D%95%E8%B0%83%E6%A0%88.md


https://leetcode.com/discuss/interview-question/352460/Google-Online-Assessment-Questions


Patience sorting
https://www.cs.princeton.edu/courses/archive/spring13/cos423/lectures/LongestIncreasingSubsequence.pdf
https://en.wikipedia.org/wiki/Patience_sorting


Trapping Rain Water（接雨水）
题目	问题	适合练习
42 https://leetcode.com/problems/trapping-rain-water/description/. 接雨水	一维柱状图，计算总积水量	前后缀最大值、双指针、单调栈
407. 接雨水 II	二维高度矩阵，计算总积水量	最小堆＋网格搜索
