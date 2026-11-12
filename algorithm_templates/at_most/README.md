# atMost 相减：恰好 K 的连续区间计数

[返回模板目录](../README.md)

## 目录

- [适用条件与证明](#invariant)
- [Go 通用骨架](#code)
- [例子与题目映射](#mapping)
- [复杂度与陷阱](#pitfalls)

<a id="invariant"></a>

## 1. 适用条件与证明

若每个区间有整数属性 count，则集合“count<=K-1”包含于“count<=K”，所以：

```text
exactly(K) = atMost(K) - atMost(K-1)
```

集合相减公式本身很通用，但**用滑窗计算 atMost**还需要：删除左端元素不增加该属性；固定 right 时，合法左端点形成连续后缀。

收缩到最早合法 left 后，所有 `[left,right]`、`[left+1,right]`……`[right,right]` 都合法，数量为 `right-left+1`，并且更早起点都非法。

不同整数种数、奇数数量满足条件；含负数数组的和通常不满足。总数量最多 n(n+1)/2，使用 int64。

<a id="code"></a>

## 2. Go 通用骨架

```go
package template

// add/remove 用数组下标更新窗口；valid 必须满足收缩可修复。
// 调用前状态为空，结束后状态可能非空；每次调用必须创建新状态。
func CountWindows(n int, add, remove func(int), valid func() bool) int64 {
    left := 0
    var answer int64
    for right := 0; right < n; right++ {
        add(right)
        for left <= right && !valid() {
            remove(left)
            left++
        }
        answer += int64(right-left+1)
    }
    return answer
}

func AtMostDistinct(nums []int, k int) int64 {
    if k < 0 { return 0 }
    freq := make(map[int]int)
    distinct := 0
    return CountWindows(len(nums),
        func(i int) {
            if freq[nums[i]] == 0 { distinct++ }
            freq[nums[i]]++
        },
        func(i int) {
            freq[nums[i]]--
            if freq[nums[i]] == 0 {
                distinct--
                delete(freq, nums[i])
            }
        },
        func() bool { return distinct <= k },
    )
}

func ExactlyDistinct(nums []int, k int) int64 {
    return AtMostDistinct(nums,k)-AtMostDistinct(nums,k-1)
}

func AtMostOdd(nums []int, k int) int64 {
    if k < 0 { return 0 }
    odd := 0
    contribution := func(i int) int {
        if nums[i]%2 != 0 { return 1 }
        return 0
    }
    return CountWindows(len(nums),
        func(i int) { odd += contribution(i) },
        func(i int) { odd -= contribution(i) },
        func() bool { return odd <= k },
    )
}

func ExactlyOdd(nums []int, k int) int64 {
    return AtMostOdd(nums,k)-AtMostOdd(nums,k-1)
}
```

<a id="mapping"></a>

## 3. 例子与题目映射

[992](https://leetcode.com/problems/subarrays-with-k-different-integers/)：nums=[1,2,1]，K=2。

```text
atMost(2) = 6：所有非空子数组
atMost(1) = 3：[1]、[2]、[1]
exactly(2) = 6-3 = 3：[1,2]、[2,1]、[1,2,1]
```

计算 atMost(2) 时，三个右端点分别贡献 1、2、3 个；每个子数组按唯一右端点计数，既不重复也不遗漏。

| 题目 | 属性 / 状态 | atMost 的合法条件 | 变化点 |
| --- | --- | --- | --- |
| [992. K 个不同整数的子数组](https://leetcode.com/problems/subarrays-with-k-different-integers/) | 频次表＋distinct | distinct<=K | 频次从 0→1 和 1→0 才修改种数 |
| [1248. 统计优美子数组](https://leetcode.com/problems/count-number-of-nice-subarrays/) | 奇数数量 | odd<=K | 奇数贡献 1、偶数贡献 0 |
| [930. 和相同的二元子数组](https://leetcode.com/problems/binary-subarrays-with-sum/) | 0/1 窗口和 | sum<=goal | 非负确保单调，goal=0 也适用 |

1248 也可以用奇数前缀数量＋哈希表：查询 prefix-K。公式相同但计数机制不同，可比较 [前缀状态模板](../prefix_state_hash/README.md)。

<a id="pitfalls"></a>

## 4. 复杂度与陷阱

- 两次线性扫描仍是 O(n)。distinct 的 map 空间 O(min(n,K+1))；odd 状态空间 O(1)。
- k<0 必须先返回 0，否则窗口清空也无法满足条件。恰好 0 时会调用 atMost(-1)。
- 每次 atMost 都初始化新的 left、freq、answer，不能沿用第一次的窗口状态。
- map 删除零频键可用 len(freq) 表示种数；若不删除，必须另存 distinct，不能混用。
- `right-left+1` 是非空窗口个数；窗口为空时为 0，不要额外计入空子数组。
- 统计的是区间数量，不是最长窗口：应累加当前所有合法后缀，而非每个 right 只加 1。
- 不能直接用 sum<=K 套在任意负数数组上。集合差恒等式仍成立，但滑窗计算 atMost 的方法会错。

已有练习：[992](../../sliding_window/l992/)、[其他滑窗题](../../sliding_window/)。
