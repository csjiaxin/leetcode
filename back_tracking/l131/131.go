package l131

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

func partition(s string) [][]string {
	result := [][]string{}
	path := []string{}
	starts := []int{0}
	searchChoices(func() bool {
		if starts[len(starts)-1] == len(s) {
			result = append(result, append([]string{}, path...))
		}
		return false
	},
		func() []int {
			ends := []int{}
			start := starts[len(starts)-1]
			for end := start + 1; end <= len(s); end++ {
				valid := true
				for l, r := start, end-1; l < r; l, r = l+1, r-1 {
					if s[l] != s[r] {
						valid = false
						break
					}
				}
				if valid {
					ends = append(ends, end)
				}
			}
			return ends
		}, func(end int) {
			start := starts[len(starts)-1]
			path = append(path, s[start:end])
			starts = append(starts, end)
		},
		func(end int) { path = path[:len(path)-1]; starts = starts[:len(starts)-1] })
	return result
}
