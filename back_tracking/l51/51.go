package l51

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

func solveNQueens(n int) [][]string {
	result := [][]string{}
	if n < 0 {
		return result
	}
	if n == 0 {
		return [][]string{{}}
	}
	columns := make([]bool, n)
	down, up := make([]bool, 2*n-1), make([]bool, 2*n-1)
	path := []int{}
	searchChoices(func() bool {
		if len(path) == n {
			board := make([]string, n)
			for row, col := range path {
				line := make([]byte, n)
				for i := range line {
					line[i] = '.'
				}
				line[col] = 'Q'
				board[row] = string(line)
			}
			result = append(result, board)
		}
		return false
	}, func() []int {
		choices := []int{}
		row := len(path)
		if row == n {
			return choices
		}
		for col := 0; col < n; col++ {
			if !columns[col] && !down[row-col+n-1] && !up[row+col] {
				choices = append(choices, col)
			}
		}
		return choices
	}, func(col int) {
		row := len(path)
		columns[col] = true
		down[row-col+n-1] = true
		up[row+col] = true
		path = append(path, col)
	},
		func(col int) {
			row := len(path) - 1
			columns[col] = false
			down[row-col+n-1] = false
			up[row+col] = false
			path = path[:row]
		})
	return result
}
