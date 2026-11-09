package l90

import (
	"math"
	"slices"
)

// subsetsWithDup returns all unique subsets of nums, which may contain duplicates.
func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)
	ret := make([][]int, 0)
	path := []int{}

	var dfs func(start int)
	dfs = func(start int) {
		ret = append(ret, slices.Clone(path))

		lastvisit := math.MinInt
		for i := start; i < len(nums); i++ {
			if nums[i] != lastvisit {
				path = append(path, nums[i])
				dfs(i + 1)
				path = path[:len(path)-1]
				lastvisit = nums[i]
			}
		}
	}
	dfs(0)
	return ret
}
