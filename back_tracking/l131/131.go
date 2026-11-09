package l131

import "slices"

// partition returns all partitions of s in which every substring is a palindrome.
func partition(s string) [][]string {
	L := len(s)

	ret := make([][]string, 0)

	path := make([]string, 0, L)

	var dfs func(start int)
	dfs = func(start int) {
		if start >= L {
			ret = append(ret, slices.Clone(path))
			return
		}
		for end := start; end < L; end++ {
			i, j := start, end
			for i <= j && s[i] == s[j] {
				i += 1
				j -= 1
			}
			if i > j {
				path = append(path, s[start:end+1])
				dfs(end + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return ret
}
