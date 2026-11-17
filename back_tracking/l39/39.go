package l39

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

func combinationSum(candidates []int, target int) [][]int {
	nums := append([]int{}, candidates...)
	sort.Ints(nums)
	result := [][]int{}
	path := []int{}
	starts := []int{0}
	sum := 0
	searchChoices(func() bool {
		if sum == target {
			result = append(result, append([]int{}, path...))
		}
		return false
	},
		func() []int {
			choices := []int{}
			if sum >= target {
				return choices
			}
			for i := starts[len(starts)-1]; i < len(nums) && sum+nums[i] <= target; i++ {
				choices = append(choices, i)
			}
			return choices
		}, func(i int) { path = append(path, nums[i]); sum += nums[i]; starts = append(starts, i) },
		func(i int) { path = path[:len(path)-1]; sum -= nums[i]; starts = starts[:len(starts)-1] })
	return result
}
