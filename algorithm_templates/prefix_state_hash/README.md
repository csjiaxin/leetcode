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

**通用的是处理步骤，不必把它们封装成回调函数。** 先掌握下面两种直接写法：求数量用 freq，求最长用 first。

### 2.1 求数量：保存历史状态出现次数

以 [560. 和为 K 的子数组](https://leetcode.com/problems/subarray-sum-equals-k/) 为基础模板：

```go
package template

func SubarraySum(nums []int, k int64) int64 {
    freq := map[int64]int64{0: 1} // 空前缀和为0，出现1次
    var prefix, answer int64

    for _, x := range nums {
        prefix += int64(x)         // 1. 更新当前状态
        answer += freq[prefix-k]   // 2. 查询历史状态
        freq[prefix]++             // 3. 记录当前状态
    }
    return answer
}
```

只需理解这个等式：`当前前缀和 - 历史前缀和 = k`，所以查询 `prefix-k`。不需要 step、need 等辅助函数。

### 2.2 换成余数：循环骨架不变

[974. 和可被 K 整除的子数组](https://leetcode.com/problems/subarray-sums-divisible-by-k/) 只需把状态换为前缀和的余数，查询相同余数：

```go
package template

// 题设 k>0，累计前缀和使用 int64。
func DivisibleSubarrays(nums []int, k int64) int64 {
    freq := map[int64]int64{0: 1}
    var prefix, answer int64

    for _, x := range nums {
        prefix += int64(x)
        remainder := prefix % k
        if remainder < 0 {
            remainder += k // Go 的负余数需要规范化
        }
        answer += freq[remainder]
        freq[remainder]++
    }
    return answer
}
```

两个前缀余数相同，差值就能被 k 整除。不需要额外维护复杂的模加法；本模板假设前缀总和在 int64 范围内，覆盖本题约束。

### 2.3 求最长：保存历史状态最早下标

以 [525. 连续数组](https://leetcode.com/problems/contiguous-array/) 为模板，0 记为 -1，1 记为 +1；相同前缀差之间的区间含有相同数量的 0 和 1。

```go
package template

func FindMaxLength(nums []int) int {
    first := map[int]int{0: -1} // 空前缀位于下标-1
    balance, answer := 0, 0

    for i, x := range nums {
        if x == 0 {
            balance--
        } else {
            balance++
        }

        if j, ok := first[balance]; ok {
            answer = max(answer, i-j)
        } else {
            first[balance] = i // 只记录第一次出现的位置
        }
    }
    return answer
}
```

**只记两种循环结构：**

- 求数量：更新状态 → `answer += freq[需要的历史状态]` → `freq[当前状态]++`。
- 求最长：更新状态 → 若历史状态出现过，用 `i-最早下标` 更新答案；否则记录当前下标。

以上三个代码块各自可以单独复制编译；合并到同一个文件时只保留一次 `package template`。

<a id="mapping"></a>

## 3. 推演与题目变化点

[560. 和为 K 的子数组](https://leetcode.com/problems/subarray-sum-equals-k/)：nums=[1,-1,1]、K=1。

| 当前元素 | 前缀和 | 查询历史前缀 | 查询到的次数 | 累计答案 |
| --- | --- | --- | --- | --- |
| 1 | 1 | 0 | 1（空前缀） | 1 |
| -1 | 0 | -1 | 0 | 1 |
| 1 | 1 | 0 | 2（空前缀、前两个元素） | 3 |

这里前缀和 0 出现两次，所以最后一次查询一次增加 2，不是只增加 1。

| 题目 | 当前前缀状态 | 查询历史 | 表保存什么 | 初始化 |
| --- | --- | --- | --- | --- |
| [560](https://leetcode.com/problems/subarray-sum-equals-k/) | 前缀和 | state-K | 出现次数 | freq[0]=1 |
| [974](https://leetcode.com/problems/subarray-sums-divisible-by-k/) | 规范化的余数 | 相同余数 | 出现次数 | freq[0]=1 |
| [525](https://leetcode.com/problems/contiguous-array/) | 1 数量减 0 数量 | 相同差值 | 最早下标 | first[0]=-1 |
| [1371](https://leetcode.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/) | a/e/i/o/u 各占一位，遇到就异或对应位 | 相同掩码 | 最早下标 | first[0]=-1 |

525 直接使用上面的 FindMaxLength：遇到 0 状态减 1，遇到 1 状态加 1。1371 沿用“保存最早下标”的循环，把 balance 换成五位奇偶掩码；按 ASCII 字节遍历，只有五个元音会翻位，其他字符不改变状态，最多 32 种状态。

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
