package l1031

// PrefixBest 先用严格更早的候选查询，再把当前候选加入历史最优集合。
func PrefixBest(n int, initial int, evaluate func(i, best int), candidate func(i int) int) {
	best := initial
	for i := 0; i < n; i++ {
		evaluate(i, best)
		best = max(best, candidate(i))
	}
}

// https://leetcode.com/problems/maximum-sum-of-two-non-overlapping-subarrays/
// 枚举右窗口起点，历史最优左窗口保证不重叠；两种先后顺序都计算。
func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	prefix := make([]int, len(nums)+1)
	for i, value := range nums {
		prefix[i+1] = prefix[i] + value
	}
	ordered := func(leftLen, rightLen int) int {
		count := len(nums) - leftLen - rightLen + 1
		if count <= 0 {
			return 0
		}
		answer := 0
		initial := prefix[leftLen]
		PrefixBest(count, initial, func(i, best int) {
			rightStart := leftLen + i
			answer = max(answer, best+prefix[rightStart+rightLen]-prefix[rightStart])
		}, func(i int) int {
			end := leftLen + i + 1
			if end > len(nums)-rightLen {
				return 0
			}
			return prefix[end] - prefix[end-leftLen]
		})
		return answer
	}
	return max(ordered(firstLen, secondLen), ordered(secondLen, firstLen))
}
