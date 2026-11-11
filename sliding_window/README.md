# 滑动窗口（Sliding Window）

## 1. 什么时候使用？

滑动窗口维护一个连续区间，通过右端加入、左端移除来增量更新状态，避免反复扫描整个区间。

常见题目特征：连续子数组或子串、固定长度、最多 k 次操作、最短覆盖、满足条件的区间数量。

使用前要判断：移动边界后，条件是否具有足够的单调性？例如正数数组中扩展窗口使和不减，收缩使和不增；包含负数时，209 的普通窗口逻辑不再适用。

本文统一使用闭区间 `[left, right]`，窗口长度是 `right-left+1`。每个元素最多进入、离开一次；若状态更新和检查都是 O(1)，总时间为 O(n)。

## 2. 定长窗口

适用：长度固定为 k 的区间和、字符频次匹配。

```go
left := 0
for right := 0; right < len(nums); right++ {
    // 加入 nums[right]。
    if right-left+1 > k {
        // 移除 nums[left]。
        left++
    }
    if right-left+1 == k {
        // 更新答案或检查是否匹配。
    }
}
```

| 题目 | 状态与关键点 |
| --- | --- |
| [643. Maximum Average Subarray I](https://leetcode.com/problems/maximum-average-subarray-i/) | 维护长度 k 的窗口和，取最大值后除以 k；最大和不能初始化为 0，因为可能全为负数 |
| [567. Permutation in String](https://leetcode.com/problems/permutation-in-string/) | 窗口长度为 len(s1)，比较 26 个小写字母的频次；只需判断是否存在 |
| [438. Find All Anagrams in a String](https://leetcode.com/problems/find-all-anagrams-in-a-string/) | 与 567 相同，但记录所有起点；重叠窗口也要保留 |

对于 567、438，比较 26 项频次是 O(26n)，在固定字母表下为 O(n)，辅助空间 O(26)。不需要生成字符串的全部排列。

## 3. 不定长窗口：最长合法区间

模式：先扩展，条件不合法时持续收缩，恢复合法后更新最大长度。

```go
left, best := 0, 0
for right := 0; right < len(nums); right++ {
    // 加入 nums[right]。
    for /* 窗口不合法 */ {
        // 移除 nums[left]。
        left++
    }
    best = max(best, right-left+1)
}
```

| 题目 | 合法条件 | 常用状态 |
| --- | --- | --- |
| [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/) | 没有重复字符 | 字符频次或最后出现位置 |
| [1004. Max Consecutive Ones III](https://leetcode.com/problems/max-consecutive-ones-iii/) | 窗口中的 0 不超过 k 个 | 零的数量 |

相关进阶：[904. Fruit Into Baskets](https://leetcode.com/problems/fruit-into-baskets/) 维护最多两种元素；[424. Longest Repeating Character Replacement](https://leetcode.com/problems/longest-repeating-character-replacement/) 使用“窗口长度减去最高字符频次不超过 k”的条件。424 中使用历史最大频次的优化需要单独理解其不变量。

## 4. 不定长窗口：最短满足区间

模式：扩展到满足条件后，持续收缩，直到不满足。答案在移除左端之前更新。

```go
left := 0
for right := 0; right < len(nums); right++ {
    // 加入 nums[right]。
    for /* 窗口满足要求 */ {
        // 更新最短答案。
        // 移除 nums[left]。
        left++
    }
}
```

| 题目 | 满足条件 | 注意点 |
| --- | --- | --- |
| [209. Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum/) | 正数窗口和 >= target | 无解返回 0；依赖所有元素为正 |
| [76. Minimum Window Substring](https://leetcode.com/problems/minimum-window-substring/) | 覆盖 t 的全部字符及其出现次数 | 重复字符不能只按种类判断；区分大小写；无解返回空字符串 |

76 可以维护每种字符的需求和窗口计数，再维护已匹配的字符数量或已达标的种类数，避免每次重新扫描全部需求。

## 5. 统计满足条件的子数组数量

### 5.1 最多 / 小于：固定右端点计数

收缩后，若 `[left,right]` 合法，且删除左侧元素不会破坏合法性，则所有以 right 结尾、起点在 left..right 之间的窗口都合法：

```go
answer += right - left + 1
```

- [713. Subarray Product Less Than K](https://leetcode.com/problems/subarray-product-less-than-k/)：元素均为正整数，维护乘积并在 `product >= k` 时收缩；`k <= 1` 时直接返回 0。
- [992. Subarrays with K Different Integers](https://leetcode.com/problems/subarrays-with-k-different-integers/)：使用频次表统计不同元素数量，并转换为两个“最多”计数之差。

### 5.2 恰好 k：两个最多相减

```text
exactly(k) = atMost(k) - atMost(k-1)
```

`atMost(k)` 内部使用最长合法窗口的收缩方式，但累加 `right-left+1`，而不是只更新最大长度。边界 `atMost(0)` 应正确返回 0；通用实现可对负 k 直接返回 0。

相关题：[930. Binary Subarrays With Sum](https://leetcode.com/problems/binary-subarrays-with-sum/)、[1248. Count Number of Nice Subarrays](https://leetcode.com/problems/count-number-of-nice-subarrays/)。这类差分要求统计量是整数且窗口条件具有相应单调性，不能直接套用于任意含负数的数组和问题。

### 5.3 至少覆盖：固定左端点计数

- [1358. Number of Substrings Containing All Three Characters](https://leetcode.com/problems/number-of-substrings-containing-all-three-characters/)：字符串只含 a、b、c，要求三者都出现。

当 `[left,right]` 已覆盖 a、b、c，保持 left 不变，右端点继续向右都有效，所以一次加入：

```go
answer += len(s) - right
```

然后移除左端并继续收缩。每次计入的是当前 left 对应的全部合法右端点，不会重复计数。另一种写法是收缩到不合法后，对当前 right 累加有效起点数量 left；两种计数方式不要混用。

## 6. 滑动窗口 + 单调队列

[239. Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/)：普通窗口不能在移除最大值后 O(1) 找到新最大值，因此维护值单调递减的下标队列。

1. 队首移除已经离开窗口的下标。
2. 队尾移除不大于当前值的旧下标，它们不可能比当前值更优且会更早过期。
3. 当前下标入队。
4. 窗口达到长度 k 后，队首就是最大值。

时间 O(n)，辅助空间 O(k)。已有实现见 `../monotonic_stack/l239/`，本轮不重复生成模板。

## 7. 常见错误

- 混淆“满足时收缩”和“不合法时收缩”：最短覆盖通常是前者，最长合法窗口通常是后者。
- 需要持续修复窗口时误用 if 而非 for；某些单步收缩优化有额外不变量，不要直接混用。
- 移除字符时忘记减少计数，或频次降为 0 时忘记减少不同字符数。
- 76 只检查字符种类，忽略重复需求；567、438 忘记固定窗口长度。
- 数组全负时把最大窗口和初始化成 0。
- 把正数窗口和的算法直接用于含负数的输入。
- Go 的字符串下标是字节；这些字符串题的题目字符集允许按字节处理，泛化到 Unicode 时需先明确 rune 与返回下标的语义。
- 子数组总数可能达到 n(n+1)/2；按题目约束选择计数类型，通用大输入建议使用 int64。

## 8. 目录与练习路线

本轮核心练习均有 `l题号/题号.go` 和 `l题号/题号_test.go`，函数模板使用 `panic("not implemented")`，测试留待实现后运行。

```text
l643 → l567 → l438             定长窗口
l3 → l1004                    最长合法窗口
l209 → l76                    最短满足窗口
l713 → l1358 → l992            子数组计数
```

原有 `1004.go`、`1358.go`、`567.go` 保留不动。它们各有独立的 main 及重复辅助函数，不适合直接一起编译；练习时单独测试 `l题号/` 目录。`567` 和 `438` 在 permutation 目录也有模板，这里按滑动窗口专题独立保留。

无根目录 go.mod 时，可从仓库根目录按文件运行，例如：

```bash
go test sliding_window/l643/*.go
# 仅检查模板和测试是否可编译，不执行测试：
go test -run '^$' sliding_window/l643/*.go
```

## 9. 参考资料

- [Sliding Window for Beginners: Problems, Template, Sample Solutions](https://leetcode.com/discuss/general-discussion/657507/Sliding-Window-for-Beginners-Problems-or-Template-or-Sample-Solutions)
- [How to Solve Sliding Window Problems](https://medium.com/outco/how-to-solve-sliding-window-problems-28d67601a66)
