package l39

import "slices"

// combinationSum returns all unique combinations of distinct positive candidates
// that sum to target. Each candidate may be chosen an unlimited number of times.
func combinationSum(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	ret := make([][]int, 0)
	path := []int{}

	var dfs func(left int, start int)
	dfs = func(left int, start int) {
		if left == 0 {
			ret = append(ret, slices.Clone(path))
			return
		}
		for i := start; i < len(candidates) && candidates[i] <= left; i++ {
			path = append(path, candidates[i])
			dfs(left-candidates[i], i)
			path = path[:len(path)-1]
		}
	}
	dfs(target, 0)
	return ret
}
