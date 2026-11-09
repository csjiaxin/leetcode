package l46

import "slices"

// permute returns all permutations of nums, which contains distinct integers.
func permute(nums []int) [][]int {
	L := len(nums)
	ret := make([][]int, 0)
	visit := make([]bool, L)

	path := make([]int, L)
	var dfs func(i int)
	dfs = func(i int) {
		if i >= L {
			ret = append(ret, slices.Clone(path))
			return
		}

		for j, n := range nums {
			if visit[j] == false {
				path[i] = n
				visit[j] = true
				dfs(i + 1)
				visit[j] = false
			}
		}

	}
	dfs(0)
	return ret
}
