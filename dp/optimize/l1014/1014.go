package l1014

import "math"

// https://leetcode.com/problems/best-sightseeing-pair/
// Time: O(n). Extra space: O(1).
func maxScoreSightseeingPair(values []int) int {
	L := len(values)
	valPlusI := make([]int, L)

	for i, v := range values {
		valPlusI[i] = v + i
	}
	ret := math.MinInt
	maxValPlusI := valPlusI[0]
	for j := 1; j < L; j++ {
		val := maxValPlusI + values[j] - j
		maxValPlusI = max(maxValPlusI, valPlusI[j])
		ret = max(ret, val)
	}
	return ret
}
