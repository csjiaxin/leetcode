package l46

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

func permute(nums []int) [][]int {
	result := [][]int{}
	path := []int{}
	used := make([]bool, len(nums))
	searchChoices(func() bool {
		if len(path) == len(nums) {
			result = append(result, append([]int{}, path...))
		}
		return false
	}, func() []int {
		candidates := []int{}
		for i := range nums {
			if !used[i] {
				candidates = append(candidates, i)
			}
		}
		return candidates
	}, func(i int) { used[i] = true; path = append(path, nums[i]) },
		func(i int) { used[i] = false; path = path[:len(path)-1] })
	return result
}
