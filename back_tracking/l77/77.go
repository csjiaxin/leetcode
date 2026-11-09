package l77

import "slices"

// combine returns all k-element combinations of the integers from 1 through n.
func combine(n int, k int) [][]int {
	path := make([]int, 0)
	ret := make([][]int, 0)
	if k < 0 || k > n {
		return ret
	}

	var dfs func(start int)
	dfs = func(start int) {
		if len(path) == k {
			ret = append(ret, slices.Clone(path))
			return
		}
		for i := start; i <= (n + 1 - k + len(path)); i++ {
			path = append(path, i)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(1)
	return ret
}
