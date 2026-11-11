# 排列与相关问题

## 1. 全排列：回溯

逐个位置选择尚未使用的元素，递归完成排列，再撤销选择。

- [46. Permutations](https://leetcode.com/problems/permutations/)：无重复元素的全排列，使用 `used[]` 记录已选元素，是 47 的基础版。
- [52. N-Queens II](https://leetcode.com/problems/n-queens-ii/)：51 的计数版，不需要构造棋盘。
- [526. Beautiful Arrangement](https://leetcode.com/problems/beautiful-arrangement/)：在位置 `pos` 放置未使用的 `v`，要求 `v % pos == 0` 或 `pos % v == 0`；可以回溯剪枝，也可以使用状态压缩 DP。
- [784. Letter Case Permutation](https://leetcode.com/problems/letter-case-permutation/)：逐位置枚举字母的大小写，数字直接保留；本质是分支选择，而非交换元素位置。

- [47. Permutations II](https://leetcode.com/problems/permutations-ii/)：包含重复元素的全排列。先排序，再跳过同一递归层中重复的选择，避免产生重复排列。
- [51. N-Queens](https://leetcode.com/problems/n-queens/)：相关回溯问题。逐行选择皇后所在列，并检查列和两条对角线是否冲突。

## 2. 下一个排列：Next Permutation

目标：保持元素不变，得到字典序中紧邻当前排列的下一个排列。

步骤：

1. 从右向左找到第一个满足 `nums[i] < nums[i+1]` 的位置 `i`。
2. 此时 `i` 右侧的后缀单调不增。从右向左找到第一个满足 `nums[j] > nums[i]` 的位置 `j`。
3. 交换 `nums[i]` 和 `nums[j]`，让当前位置以最小幅度变大。
4. 反转 `i` 右侧的后缀，使其按升序排列，得到最小的后缀。

如果找不到位置 `i`，说明当前已经是最大排列，反转整个数组即可得到最小排列。

时间复杂度为 O(n)，额外空间复杂度为 O(1)。

### 相关题目

- [31. Next Permutation](https://leetcode.com/problems/next-permutation/)：原地生成下一个排列。
  - [参考讨论：基于 Wikipedia 的实现](https://leetcode.com/problems/next-permutation/discuss/13867/C%2B%2B-from-Wikipedia)
- [556. Next Greater Element III](https://leetcode.com/problems/next-greater-element-iii/)：对整数的各位数字求下一个排列；如果不存在更大的排列，或结果超过 32 位有符号整数上限，返回 `-1`。
- [1850. Minimum Adjacent Swaps to Reach the Kth Smallest Number](https://leetcode.com/problems/minimum-adjacent-swaps-to-reach-the-kth-smallest-number/)：先执行 k 次下一个排列得到目标，再通过逐位匹配和相邻交换计算最少交换次数。

## 3. 下一个更大元素：单调栈

- [503. Next Greater Element II](https://leetcode.com/problems/next-greater-element-ii/)：在循环数组中，为每个元素寻找右侧第一个严格更大的元素。

注意：503 不改变数组的排列，与 31、556 的“下一个排列”不是同一类问题。它通常使用单调栈，并通过遍历两遍数组处理循环关系，时间复杂度为 O(n)，额外空间复杂度为 O(n)。

## 4. 字符串中的排列：滑动窗口

- [567. Permutation in String](https://leetcode.com/problems/permutation-in-string/)：判断 `s2` 是否包含一个子串，其字符频次与 `s1` 完全相同。

不需要枚举 `s1` 的全部排列。维护长度为 `len(s1)` 的滑动窗口，比较字符频次即可。对于题目限定的小写英文字母，时间复杂度为 O(len(s1) + len(s2))，辅助空间为 O(26)。

- [438. Find All Anagrams in a String](https://leetcode.com/problems/find-all-anagrams-in-a-string/)：与 567 使用相同的定长窗口，但返回全部匹配起点；匹配窗口可以重叠。
- [242. Valid Anagram](https://leetcode.com/problems/valid-anagram/)：只比较两个完整字符串的字符频次。
- [767. Reorganize String](https://leetcode.com/problems/reorganize-string/)：重排后相邻字符不能相同，使用最大堆或按位置贪心填充；不需要枚举全部排列。

## 5. 字典序进阶：上一个排列、排名与第 k 个排列

### 1053. 上一个排列：仅允许一次交换

[Previous Permutation With One Swap](https://leetcode.com/problems/previous-permutation-with-one-swap/)

1. 从右向左找到第一个满足 `arr[i] > arr[i+1]` 的位置。
2. 在右侧找到严格小于 `arr[i]` 的最大值。
3. 若该值重复，选择其中最靠左的位置交换，例如 `[3,1,1,3] → [1,3,1,3]`。

只能交换一次，因此不能像 31 一样反转后缀。不存在更小排列时返回原排列。

### 60. 第 k 个排列：阶乘分组

[Permutation Sequence](https://leetcode.com/problems/permutation-sequence/)

`1..n` 的排列中，每个首位数字对应 `(n-1)!` 个排列。先将 `k` 减一，转换为从 0 开始的编号，再逐位选择：

```text
index = k / (remaining-1)!
k     = k % (remaining-1)!
```

取剩余有序数字中的第 `index` 个并移除。不必执行 `k-1` 次下一个排列；使用切片删除时，总时间 O(n²)，辅助空间 O(n)。

### 1830. 排列排名：重复字符的组合计数

[Minimum Number of Operations to Make String Sorted](https://leetcode.com/problems/minimum-number-of-operations-to-make-string-sorted/)

逐位统计当前位置换成更小字符时，有多少个不同的后续排列。若剩余字符数为 `m`，各字符频次为 `c1, c2, ...`，排列数为 `m! / (c1! × c2! × ...)`。题目需要模运算，可预处理阶乘和阶乘逆元。

## 6. 根据约束构造排列：贪心与双指针

- [942. DI String Match](https://leetcode.com/problems/di-string-match/)：排列使用 `0..n`。维护未使用的最小值 `low` 和最大值 `high`；遇到 `I` 取 `low`，遇到 `D` 取 `high`，最后填入剩余值。返回任意合法排列即可，时间 O(n)。
- [2375. Construct Smallest Number From DI String](https://leetcode.com/problems/construct-smallest-number-from-di-string/)：使用不重复的数字 `1..9`，要求结果最小。数字递增入栈，遇到 `I` 或末尾时统一弹出，反转连续 `D` 对应的片段。例如 `IIDDD → 126543`，时间 O(n)。
- [667. Beautiful Arrangement II](https://leetcode.com/problems/beautiful-arrangement-ii/)：要求恰好 `k` 种相邻绝对差。在 `1..k+1` 内交替取最小和最大值，产生差值 `k,k-1,...,1`，然后升序接上剩余数字。

## 7. 排列计数：动态规划

### 903. Valid Permutations for DI Sequence

[题目链接](https://leetcode.com/problems/valid-permutations-for-di-sequence/)

与 942 类似，但要求统计所有方案。令 `dp[i][j]` 表示使用 `i+1` 个数，末尾数字在这些数中的排名为 `j`，满足前 `i` 个关系的方案数。遇到 `I` 累加前一层排名小于 `j` 的状态，遇到 `D` 累加排名大于等于 `j` 的状态；通过前缀和、后缀和将总时间优化到 O(n²)。

### 629. K Inverse Pairs Array

[题目链接](https://leetcode.com/problems/k-inverse-pairs-array/)

令 `dp[n][k]` 表示 `1..n` 的排列中恰好有 `k` 个逆序对的数量。将最大值 `n` 插入长度 `n-1` 的排列，会新增 `0..n-1` 个逆序对：

```text
dp[n][k] = sum(dp[n-1][k-x]), 0 <= x <= min(k,n-1)
```

通过滑动窗口优化到 O(nk) 时间，滚动数组空间为 O(k)。

## 8. 排列转换：交换与逆序对

- [765. Couples Holding Hands](https://leetcode.com/problems/couples-holding-hands/)：通过任意位置交换，让每对伴侣相邻。可使用贪心加位置表，或并查集。
- [854. K-Similar Strings](https://leetcode.com/problems/k-similar-strings/)：BFS 求两个字母异位词之间的最少任意位置交换；优先修复第一个不匹配位置来剪枝。
- [775. Global and Local Inversions](https://leetcode.com/problems/global-and-local-inversions/)：判断是否存在非相邻逆序对。对 `0..n-1` 的排列，所有位置都满足 `abs(nums[i]-i) <= 1` 当且仅当全局逆序对数等于局部逆序对数。

注意区分交换规则：1850 只允许相邻交换，854 允许任意位置交换，765 的目标是伴侣相邻而非固定排列。

## 9. 推荐练习路线

```text
46 → 47 → 51 → 52 → 526       回溯与计数
31 → 556 → 1053 → 60          字典序排列
942 → 2375 → 903              构造一个解与统计全部解
567 → 438                     字符频次滑动窗口
1850 → 854                    不同交换规则
```

优先补充的七道题：46、60、1053、526、942、2375、438。对应目录采用 `l题号/题号.go` 和 `l题号/题号_test.go`，函数模板留待实现。