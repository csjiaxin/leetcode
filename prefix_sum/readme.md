# 前缀状态与边界增量

本目录合并原 `Subarray` 中的前缀状态题与原 `difference_array`。两个子系列共享“用前缀恢复区间信息”的来源，各自保留明确的状态和操作模板。

| 子系列 | 适用条件 | 固定不变量 | 模板与题目 |
| --- | --- | --- | --- |
| [前缀状态＋哈希](state_hash/readme.md) | 区间条件可以写成两个前缀状态的关系 | 查询时 map 只包含当前位置以前的前缀 | [Count / Longest](state_hash/template.go)，560、974、525、1371、1248 |
| [差分与扫描](difference/README.md) | 多次可叠加的区间加法，最后统一恢复 | 左端加入贡献，右端之后取消贡献 | `diff[l]+=x; diff[r]-=x`，1094、1109、1854 |

## 计数与最长长度分别复用什么

```go
// 状态推进由 advance 决定；当前前缀需要的历史状态由 needed 决定。
Count(values, initial, advance, needed)
Longest(values, initial, advance, needed)
```

`Count` 保存历史频次，先查询再增加当前频次；`Longest` 保存最早位置，第一次出现后不覆盖。初始化分别是空前缀频次 1、空前缀位置 -1。

| 题目 | 前缀状态 | 需要的历史状态 | 聚合 |
| --- | --- | --- | --- |
| [560](https://leetcode.com/problems/subarray-sum-equals-k/) | 累计和 s | s-k | Count |
| [974](https://leetcode.com/problems/subarray-sums-divisible-by-k/) | 规范化余数 r | r | Count |
| [525](https://leetcode.com/problems/contiguous-array/) | 1 的数量减 0 的数量 | 相同状态 | Longest |
| [1371](https://leetcode.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/) | 五种元音出现次数的奇偶 mask | 相同 mask | Longest |
| [1248](https://leetcode.com/problems/count-number-of-nice-subarrays/) | 累计奇数个数 c | c-k | Count |

实际可执行模板和五种适配示例由 [template_test.go](state_hash/template_test.go) 验证。五道题均已在源码中内嵌相应 Count 或 Longest 模板，README 中的适配器与实际实现一致。

## 差分的固定骨架

采用半开区间 `[l,r)`：

```go
diff := make([]int, n+1)
for _, update := range updates {
    diff[update.left] += update.delta
    diff[update.right] -= update.delta
}
current := 0
for i := 0; i < n; i++ {
    current += diff[i]
    answer[i] = original[i] + current
}
```

这里 diff 从零开始，只记录增量；如果 diff 已由原数组构造，恢复时应直接使用 current。闭区间、上下车时刻与年份边界的转换见子系列文档。

前缀哈希与差分不共用同一个 map 更新规则；动态更新与查询则另见 [线段树](../tree/segment/readme.md)。原 `Subarray` 中的单调栈、双指针草稿已按其实际算法分别移入对应题目目录。
