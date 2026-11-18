不一定减少状态数量，而是避免每个状态都重新枚举所有前驱。

[1937](https://leetcode.com/problems/maximum-number-of-points-with-cost/) 就是把 O(mn²) 的转移优化到 O(mn)。下面这些很适合一起练，复杂度按“直接枚举 → 常见优化解法”比较：
题目	优化方法	时间复杂度
[1014](https://leetcode.com/problems/best-sightseeing-pair/). 最佳观光组合	拆分公式，维护前缀最大值	O(n²) → O(n)
[1289](https://leetcode.com/problems/minimum-falling-path-sum-ii/). 下降路径最小和 II	维护上一行不同列的最小、次小值	O(n³) → O(n²)
[1031](https://leetcode.com/problems/maximum-sum-of-two-non-overlapping-subarrays/). 两个非重叠子数组的最大和	前缀和＋维护前面最佳窗口	O(n²) → O(n)
[188](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/). 买卖股票的最佳时机 IV	维护最佳买入状态，避免枚举买入日	O(kn²) → O(kn)
[1696](https://leetcode.com/problems/jump-game-vi/). 跳跃游戏 VI	单调队列维护窗口内的最大 DP 值	O(nk) → O(n)
[1425](https://leetcode.com/problems/constrained-subsequence-sum/). 带限制的子序列和	单调队列优化有限范围的转移	O(nk) → O(n)
[300](https://leetcode.com/problems/longest-increasing-subsequence/). 最长递增子序列	改用“最小结尾”状态＋二分	O(n²) → O(n log n)
[1235](https://leetcode.com/problems/maximum-profit-in-job-scheduling/). 规划兼职工作	排序＋二分查找兼容的前驱工作	O(n²) → O(n log n)
