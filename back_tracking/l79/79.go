package l79

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

type boardCell struct{ row, col int }

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	rows, cols := len(board), len(board[0])
	var frequency [256]int
	for _, row := range board {
		for _, value := range row {
			frequency[value]++
		}
	}
	for i := range word {
		frequency[word[i]]--
		if frequency[word[i]] < 0 {
			return false
		}
	}
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}
	path := []boardCell{}
	return searchChoices(func() bool { return len(path) == len(word) },
		func() []boardCell {
			choices := []boardCell{}
			add := func(row, col int) {
				if row >= 0 && row < rows && col >= 0 && col < cols && !visited[row][col] && board[row][col] == word[len(path)] {
					choices = append(choices, boardCell{row, col})
				}
			}
			if len(path) == 0 {
				for row := 0; row < rows; row++ {
					for col := 0; col < cols; col++ {
						add(row, col)
					}
				}
			} else {
				last := path[len(path)-1]
				add(last.row-1, last.col)
				add(last.row+1, last.col)
				add(last.row, last.col-1)
				add(last.row, last.col+1)
			}
			return choices
		}, func(cell boardCell) { visited[cell.row][cell.col] = true; path = append(path, cell) },
		func(cell boardCell) { visited[cell.row][cell.col] = false; path = path[:len(path)-1] })
}
