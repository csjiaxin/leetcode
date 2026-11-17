# 背包：容量、选择次数与计数顺序

本目录合并原 coin 零钱题与 distinct_ways 中真正按目标容量转移的题。物品只能选一次还是可以重复选、求最值还是方案数、是否区分顺序，决定不同模板。

## 1. 0/1 背包：每个输入元素最多使用一次

一维压缩且物品在外层时，容量倒序：

~~~go
for _, weight := range weights {
    for capacity := limit; capacity >= weight; capacity-- {
        dp[capacity] = combine(dp[capacity], dp[capacity-weight])
    }
}
~~~

这是循环骨架，combine 和初始化由状态含义决定。不变量：当前物品的转移来源仍属于处理它之前的状态，避免同一次外层循环反复选择它。

| 目标 | 空容量 | 不可达容量 | 合并规则 |
| --- | --- | --- | --- |
| 是否可达 | true | false | 逻辑或 |
| 方案数 | 1 | 0 | 加法 |
| 最少数量 | 0 | 无穷大 | 对可达前驱加 1，再取 min |
| 最大收益 | 按是否要求恰好装满初始化 | 不能一律当成 0 | 取 max |

- [416](../l416/416.go)：是否能达到总和的一半。
- [1049](../l1049/1049.go)：不超过一半的最大可达和；[旧多重集搜索](../l1049/legacy/1049)保留为另一种历史方法。
- [494](../l494/494.go)：把正负号分配化为子集计数；应先推导总和、目标范围与奇偶条件。零元素也只能选一次，会使相应计数翻倍。[旧草稿](../l494/legacy/494)仍保留。
- [474](../l474/474.go)：0 和 1 是两种资源，两个容量都倒序，每个字符串只选一次。

## 2. 完全背包：允许重复使用

物品在外层时，容量正序，使来源可以包含当前物品：

~~~go
for _, coin := range coins {
    for amount := coin; amount <= target; amount++ {
        dp[amount] = combine(dp[amount], dp[amount-coin])
    }
}
~~~

[322](../l322/322.go) 求最少硬币数。dp[0]=0，其他状态初始不可达；读取无穷大时不能直接加 1。硬币面额必须满足题目中的正数条件。

## 3. 组合计数与有序序列计数

[518](../l518/518.go) 不区分选币顺序：硬币在外层、金额正序。

~~~go
dp[0] = 1
for _, coin := range coins {
    for amount := coin; amount <= target; amount++ {
        dp[amount] += dp[amount-coin]
    }
}
~~~

[377](../l377/377_dp) 区分顺序：目标金额在外层，枚举最后选的数字。

~~~go
dp[0] = 1
for amount := 1; amount <= target; amount++ {
    for _, value := range nums {
        if value <= amount {
            dp[amount] += dp[amount-value]
        }
    }
}
~~~

例如面额 1、2 凑成 3，518 的组合为 111、12；377 的序列还区分 21。不能笼统说“所有背包只需改正序/倒序”；外层枚举对象同样影响含义。

保留的替代材料：

- [518 记忆化版本](../l518/legacy/recursive/518.go)。
- [377 递归片段](../l377/377)与[迭代片段](../l377/377_dp)。
- [原 coin 笔记](../legacy/coin.readme.md)、[原 knapsack 笔记](../legacy/knapsack.readme.md)。
- [背包问题九讲原 PDF](../%E8%83%8C%E5%8C%85%E9%97%AE%E9%A2%98%E4%B9%9D%E8%AE%B2V2.pdf)。

## 归类边界

原 931 是网格下降路径，已进入 [grid/l931](../../grid/l931/931.go)。原 871 实际使用大顶堆贪心，位于 [heap_priority_queue](../../../heap_priority_queue/l871/legacy/from_knapsack/871)。

650 与误命名为 1240 的相同复制/粘贴片段进入 [linear/copy_paste](../../linear/README.md)，983 进入 linear，343 的未完成草稿也保留为独立包；不把这些代码硬套到容量倒序模板。

一维容量为 C、物品数为 n 时，常见时间 O(nC)、辅助空间 O(C)；474 等多资源状态还要乘其他容量维度。
