package l78

import "slices"

func subsets(nums []int) [][]int {
	r := make([][]int, 0)
	r = append(r, []int{})
	for _, n := range nums {
		L := len(r)
		for i := 0; i < L; i++ {
			e := slices.Clone(r[i])
			r = append(r, append(e, n))
		}
	}
	return r
}
