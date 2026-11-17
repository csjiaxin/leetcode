package l47

import (
	"math"
	"slices"
)

// https://leetcode.com/problems/permutations-ii/
// permuteUnique returns all distinct permutations in any order.
func permuteUnique(nums []int) [][]int {
	slices.Sort(nums)
	N := len(nums)
	visit := make([]bool, N)

	path := make([]int, 0, N)
	result := make([][]int, 0)
	var dfs func()
	dfs = func() {
		if len(path) == N {
			result = append(result, slices.Clone(path))
			return
		}
		lastVisit := math.MinInt
		for i := 0; i < N; i++ {
			if visit[i] == false && nums[i] != lastVisit {
				lastVisit = nums[i]
				visit[i] = true
				path = append(path, nums[i])

				dfs()

				visit[i] = false
				path = path[:len(path)-1]
			}
		}
	}
	dfs()
	return result
}
