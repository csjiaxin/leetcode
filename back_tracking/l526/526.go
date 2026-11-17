package l526

import "sort"

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

func countArrangement(n int) int {
	// Visit the most constrained positions first; position order does not change the count.
	possible := make([][]int, n)
	for position := 1; position <= n; position++ {
		for value := 1; value <= n; value++ {
			if value%position == 0 || position%value == 0 {
				possible[position-1] = append(possible[position-1], value)
			}
		}
	}
	sort.Slice(possible, func(i, j int) bool { return len(possible[i]) < len(possible[j]) })
	used := make([]bool, n+1)
	depth, count := 0, 0
	searchChoices(func() bool {
		if depth == n {
			count++
		}
		return false
	},
		func() []int {
			choices := []int{}
			if depth == n {
				return choices
			}
			for _, value := range possible[depth] {
				if !used[value] {
					choices = append(choices, value)
				}
			}
			return choices
		}, func(value int) { used[value] = true; depth++ }, func(value int) { used[value] = false; depth-- })
	return count
}
