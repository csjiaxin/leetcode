# 不定长滑动窗口

[返回模板目录](../README.md)

## 目录

- [适用条件与不变量](#invariant)
- [Go 固定骨架](#code)
- [变化点与推演](#mapping)
- [复杂度与易错点](#pitfalls)

<a id="invariant"></a>

## 1. 适用条件与不变量

处理连续区间。必须证明窗口状态随着边界移动有足够的单调性，不能因为题目含“子数组”就使用滑窗。

**求最长合法窗口**：扩张可能违规，收缩可以修复；固定 right，合法左端点是一个后缀。收缩结束，left 是当前最早合法起点，窗口为以 right 结尾的最长合法窗口。

**求最短满足窗口**：扩张可能满足要求，持续收缩直到不满足；在收缩前记录合法答案。例如正数数组的和至少 target、字符串覆盖所有目标字符。

以下通用回调把加入、移除、判断合法性从边界骨架分离。所有回调共同维护同一份窗口状态，初始为空。

<a id="code"></a>

## 2. Go 固定骨架

```go
package template

// 适用于缩短合法窗口仍合法的问题；empty 窗口应合法。
func LongestWindow(n int, add, remove func(int), valid func() bool) int {
    left, answer := 0, 0
    for right := 0; right < n; right++ {
        add(right)
        for left <= right && !valid() {
            remove(left)
            left++
        }
        answer = max(answer, right-left+1)
    }
    return answer
}

// 适用于扩张满足窗口仍满足的问题；空窗口应不满足。
// 不存在时返回 0。
func ShortestWindow(n int, add, remove func(int), satisfied func() bool) int {
    left, answer := 0, n+1
    for right := 0; right < n; right++ {
        add(right)
        for left <= right && satisfied() {
            answer = min(answer, right-left+1)
            remove(left)
            left++
        }
    }
    if answer == n+1 { return 0 }
    return answer
}

// ASCII / 字节语义，对应 LC3 的字符范围。
func LongestUnique(s string) int {
    var freq [256]int
    duplicates := 0
    return LongestWindow(len(s),
        func(i int) {
            freq[s[i]]++
            if freq[s[i]] == 2 { duplicates++ }
        },
        func(i int) {
            if freq[s[i]] == 2 { duplicates-- }
            freq[s[i]]--
        },
        func() bool { return duplicates == 0 },
    )
}

func MinPositiveSum(nums []int, target int64) int {
    if target <= 0 { panic("target must be positive") }
    var sum int64
    return ShortestWindow(len(nums),
        func(i int) { sum += int64(nums[i]) },
        func(i int) { sum -= int64(nums[i]) },
        func() bool { return sum >= target },
    )
}
```

<a id="mapping"></a>

## 3. 变化点与推演

[3. 无重复字符的最长子串](https://leetcode.com/problems/longest-substring-without-repeating-characters/) 的 s="abba"：

| right | 加入后 | 收缩过程 | 合法窗口 | 最长长度 |
| --- | --- | --- | --- | --- |
| 0 | a | 无 | a | 1 |
| 1 | ab | 无 | ab | 2 |
| 2 | abb | 移除 a 后仍重复，再移除 b | b | 2 |
| 3 | ba | 无 | ba | 2 |

为什么必须 while 而不是 if？一次移除不一定能修复违规，abb 就是反例。

| 题目 | add/remove 状态 | 条件 | 骨架和更新时机 |
| --- | --- | --- | --- |
| [3](https://leetcode.com/problems/longest-substring-without-repeating-characters/) | 字符频次、重复字符种数 | 没有重复 | LongestWindow，修复后更新 |
| [1004](https://leetcode.com/problems/max-consecutive-ones-iii/) | 零的个数 | zeros<=k | LongestWindow |
| [209](https://leetcode.com/problems/minimum-size-subarray-sum/) | 正数窗口和 | sum>=target | ShortestWindow，每次移除前更新 |
| [76](https://leetcode.com/problems/minimum-window-substring/) | 目标频次和满足的字符种数 | 所有需求均满足 | ShortestWindow 的思路，额外保存最优起点；重数字符按次数满足 |

76 需要返回字符串，不是仅返回长度：答案变好时同时记 left，最后切片。target 为空时单独返回空串。

<a id="pitfalls"></a>

## 4. 复杂度与易错点

- 每个元素最多加入一次、移除一次，若回调 O(1)，总时间 O(n)，而非 O(n²)。频次空间取决于字符集或不同元素数。
- 含负数的和不满足普通滑窗单调性：[1,-1,5]、target=5，按正数模板会停在长度 3，却漏掉 [5]；应考虑前缀和及其他结构。
- 最长合法与最短满足的 while 条件相反，更新答案时机也不同。
- 删除字符时维护频次变化临界点，不能每次删除都减少“满足种数”。
- Unicode 字符要转 []rune，并明确返回字符长度还是字节位置，不能照搬 [256]int。
- 一般模板要求每轮完整修复；某些题存在只维护长度的非收缩变体，不要混进本模板。

已有练习：[sliding_window](../../sliding_window/)。计数见 [atMost](../at_most/README.md)，负数和计数见 [前缀状态](../prefix_state_hash/README.md)。
