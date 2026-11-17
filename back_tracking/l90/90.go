package l90

import "sort"

// Standalone choose/recurse/undo subtemplate and problem adapter.
// searchChoices restores every choice even when a successful search stops early.
// visit may record a result; returning true requests early termination.
func searchChoices[T any](visit func() bool, choices func() []T, choose, undo func(T)) bool {
	if visit() {
		return true
	}
	for _, choice := range choices() {
		choose(choice)
		found := searchChoices(visit, choices, choose, undo)
		undo(choice)
		if found {
			return true
		}
	}
	return false
}

func subsetsWithDup(nums []int) [][]int {
	nums = append([]int{}, nums...)
	sort.Ints(nums)
	result := [][]int{}
	path := []int{}
	starts := []int{0}
	searchChoices(func() bool { result = append(result, append([]int{}, path...)); return false },
		func() []int {
			candidates := []int{}
			start := starts[len(starts)-1]
			for i := start; i < len(nums); i++ {
				if i == start || nums[i] != nums[i-1] {
					candidates = append(candidates, i)
				}
			}
			return candidates
		}, func(i int) { path = append(path, nums[i]); starts = append(starts, i+1) },
		func(i int) { path = path[:len(path)-1]; starts = starts[:len(starts)-1] })
	return result
}
