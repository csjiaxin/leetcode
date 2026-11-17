package grid

import "testing"

func TestFloodFillComponentLifecycle(t *testing.T) {
	cells := [][]bool{{true, false, true}, {true, false, true}}
	visited := [][]bool{make([]bool, 3), make([]bool, 3)}
	counts := make(map[[2]int]int)
	enter := func(r, c int) bool { return cells[r][c] }
	visit := func(r, c int) {
		if !visited[r][c] {
			t.Fatal("visit must observe an already marked cell")
		}
		counts[[2]int{r, c}]++
	}
	FloodFill(2, 3, -1, 0, visited, enter, visit)
	FloodFill(2, 3, 0, 1, visited, enter, visit)
	if len(counts) != 0 {
		t.Fatal("invalid or blocked start was visited")
	}
	FloodFill(2, 3, 0, 0, visited, enter, visit)
	FloodFill(2, 3, 1, 0, visited, enter, visit)
	if len(counts) != 2 || visited[0][2] {
		t.Fatal("component boundary or shared visited was ignored")
	}
	FloodFill(2, 3, 0, 2, visited, enter, visit)
	if len(counts) != 4 {
		t.Fatal("second component was not visited")
	}
	for cell, count := range counts {
		if count != 1 {
			t.Fatalf("%v visited %d times", cell, count)
		}
	}
	if visited[0][1] || visited[1][1] {
		t.Fatal("blocked cells were marked")
	}
}

func TestFloodFillUsesFourNeighbors(t *testing.T) {
	visited := [][]bool{make([]bool, 2), make([]bool, 2)}
	count := 0
	FloodFill(2, 2, 0, 0, visited,
		func(r, c int) bool { return r == c },
		func(r, c int) { count++ },
	)
	if count != 1 || visited[1][1] {
		t.Fatal("diagonal cells must belong to separate components")
	}
}
