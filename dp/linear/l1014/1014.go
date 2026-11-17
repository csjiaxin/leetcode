package l1014

// PrefixBest 先用严格更早的候选查询，再把当前候选加入历史最优集合。
func PrefixBest(n int, initial int, evaluate func(i, best int), candidate func(i int) int) {
	best := initial
	for i := 0; i < n; i++ {
		evaluate(i, best)
		best = max(best, candidate(i))
	}
}

// https://leetcode.com/problems/best-sightseeing-pair/
// 历史项是 values[i]+i；先查询再加入，确保 i<j。
func maxScoreSightseeingPair(values []int) int {
	if len(values) < 2 {
		return 0
	}
	answer := 0
	PrefixBest(len(values)-1, values[0], func(i, best int) {
		j := i + 1
		answer = max(answer, best+values[j]-j)
	}, func(i int) int { j := i + 1; return values[j] + j })
	return answer
}
