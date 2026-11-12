# 前缀状态＋哈希表

[返回模板目录](../README.md)

## 目录

- [适用条件与不变量](#invariant)
- [通用 Go 骨架](#code)
- [推演与题目变化点](#mapping)
- [复杂度与易错点](#pitfalls)

<a id="invariant"></a>

## 1. 适用条件与不变量

连续区间 `[l,r]` 可以用两个前缀状态的关系表达：前缀和之差、余数相同、0/1 数量差相同、奇偶掩码异或为零。**允许负数**，不依赖窗口和的单调性。

以 P[r] 表示到 r 的前缀状态，空前缀 P[-1]=0：

- 和为 K：P[r]-P[l-1]=K，因此查询历史状态 P[r]-K。
- 和被 K 整除：两个前缀的规范化余数相同。
- 0/1 数量相等：把 0 记为 -1、1 记为 +1，两个前缀差相同。
- 元音全为偶数：五位奇偶掩码相同，区间异或为零。

**不变量**：查询时，哈希表只保存当前元素之前的前缀（含空前缀）；查询结束才记录当前状态，避免把空区间计入答案。

固定骨架：初始化空前缀 → 更新当前状态 → 查询历史 → 记录当前状态。

<a id="code"></a>

## 2. 通用 Go 骨架

用 step 定义状态更新、need 定义需要匹配的历史状态。状态用 int64；复杂掩码也可映射为该整数类型。

```go
package template

// CountByPrefix 统计非空连续区间；单次 step/need 应为 O(1)。
func CountByPrefix(nums []int, initial int64,
    step func(int64, int) int64,
    need func(int64) int64,
) int64 {
    freq := map[int64]int64{initial: 1}
    state := initial
    var answer int64
    for _, x := range nums {
        state = step(state, x)
        answer += freq[need(state)] // 先查过去
        freq[state]++              // 后记录现在
    }
    return answer
}

func SubarraySum(nums []int, k int64) int64 {
    return CountByPrefix(nums, 0,
        func(state int64, x int) int64 { return state + int64(x) },
        func(state int64) int64 { return state - k },
    )
}

func DivisibleSubarrays(nums []int, k int64) int64 {
    if k <= 0 { panic("k must be positive") }
    return CountByPrefix(nums, 0,
        func(state int64, x int) int64 {
            // 规范化到 [0,k)，避免先加 k 造成溢出。
            xmod := int64(x)%k
            if xmod < 0 { xmod += k }
            if state >= k-xmod { return state-(k-xmod) }
            return state+xmod
        },
        func(state int64) int64 { return state },
    )
}

// LongestByPrefix 查询相同历史状态，保留最早下标。
func LongestByPrefix(nums []int, initial int64,
    step func(int64, int) int64,
) int {
    first := map[int64]int{initial: -1}
    state, answer := initial, 0
    for i, x := range nums {
        state = step(state, x)
        if j, ok := first[state]; ok {
            answer = max(answer, i-j)
        } else {
            first[state] = i // 不覆盖最早下标
        }
    }
    return answer
}
```

这里是状态模式的抽象，不代表所有前缀题都必须强行用回调。面试中可直接展开，先写清等式，再选择计数或最早位置。

<a id="mapping"></a>

## 3. 推演与题目变化点

[560. 和为 K 的子数组](https://leetcode.com/problems/subarray-sum-equals-k/)：nums=[1,-1,1]、K=1。

| 当前元素 | 前缀和 | 查询历史前缀 | 查询到的次数 | 累计答案 |
| --- | --- | --- | --- | --- |
| 1 | 1 | 0 | 1（空前缀） | 1 |
| -1 | 0 | -1 | 0 | 1 |
| 1 | 1 | 0 | 2（空前缀、前两个元素） | 3 |

这里前缀和 0 出现两次，所以最后一次查询一次增加 2，不是只增加 1。

| 题目 | step / 状态 | 查询历史 | 表保存什么 | 初始化 |
| --- | --- | --- | --- | --- |
| [560](https://leetcode.com/problems/subarray-sum-equals-k/) | 前缀和 | state-K | 出现次数 | freq[0]=1 |
| [974](https://leetcode.com/problems/subarray-sums-divisible-by-k/) | 规范化的余数 | 相同余数 | 出现次数 | freq[0]=1 |
| [525](https://leetcode.com/problems/contiguous-array/) | 1 数量减 0 数量 | 相同差值 | 最早下标 | first[0]=-1 |
| [1371](https://leetcode.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/) | a/e/i/o/u 各占一位，遇到就异或对应位 | 相同掩码 | 最早下标 | first[0]=-1 |

525 可使用 LongestByPrefix：遇到 0 状态减 1，遇到 1 状态加 1。1371 可按 ASCII 字节遍历字符串；只有五个元音会翻位，其他字符不改变状态，最多 32 种状态。

**次数 vs 最早位置**：同一状态出现越多，计数贡献越多；求最长时则只需要最早位置，晚位置不会产生更长区间。

<a id="pitfalls"></a>

## 4. 复杂度与易错点

- 平均时间 O(n)，额外空间 O(n)；974 至多 min(n+1,K) 种余数，1371 只有 32 种状态。
- Go 的负数取模仍可能为负：先 `r=x%k`，若 r<0 则 r+=k，得到 [0,k) 的规范化余数。
- 先记录当前再查询，K=0 时会误计空区间。nums=[0] 的正确答案为 1，而不是 2。
- 求最长时覆盖 first[state] 会丢失最早边界；first 为 0 不能代表“不存在”，须使用 map 的 ok。
- 区间数量最多 n(n+1)/2，使用 int64；前缀和也应检查题目数值范围。
- 这不是通用的区间最大值查询：最大值不能由两个普通前缀最大值相减恢复。
- [1248](https://leetcode.com/problems/count-number-of-nice-subarrays/) 也可把奇数记为 1、偶数记为 0 后使用前缀计数，与 [atMost 模板](../at_most/README.md) 对照。

已有相关练习：[滑动窗口](../../sliding_window/)、[子数组专题](../../Subarray/)。
