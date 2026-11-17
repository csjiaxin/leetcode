package l77

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

func combine(n int, k int) [][]int {
	result := [][]int{}
	if k < 0 || k > n {
		return result
	}
	path := []int{}
	searchChoices(func() bool {
		if len(path) == k {
			result = append(result, append([]int{}, path...))
		}
		return false
	},
		func() []int {
			candidates := []int{}
			if len(path) == k {
				return candidates
			}
			start := 1
			if len(path) > 0 {
				start = path[len(path)-1] + 1
			}
			for value := start; value <= n-(k-len(path))+1; value++ {
				candidates = append(candidates, value)
			}
			return candidates
		}, func(value int) { path = append(path, value) }, func(value int) { path = path[:len(path)-1] })
	return result
}
