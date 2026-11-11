# 排列与相关问题

## 1. 全排列：回溯

逐个位置选择尚未使用的元素，递归完成排列，再撤销选择。

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