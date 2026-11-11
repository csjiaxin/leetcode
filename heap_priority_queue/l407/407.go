package l407

import "leetcode/heap_priority_queue/internal/pqueue"

type cell struct{ row, col, height int }

// trapRainWater expands the lowest boundary; visited cells are marked on insertion.
// O(mn log(mn)) time, O(mn) space.
func trapRainWater(heightMap [][]int) int {
	m := len(heightMap)
	if m == 0 {
		return 0
	}
	n := len(heightMap[0])
	if m < 3 || n < 3 {
		return 0
	}
	visited := make([][]bool, m)
	for r := range visited {
		visited[r] = make([]bool, n)
	}
	q := pqueue.New(func(a, b cell) bool { return a.height < b.height })
	add := func(r, c int) {
		if !visited[r][c] {
			visited[r][c] = true
			q.Push(cell{r, c, heightMap[r][c]})
		}
	}
	for r := 0; r < m; r++ {
		add(r, 0)
		add(r, n-1)
	}
	for c := 0; c < n; c++ {
		add(0, c)
		add(m-1, c)
	}
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	water := 0
	for q.Len() > 0 {
		s := q.Pop()
		for _, d := range dirs {
			r, c := s.row+d[0], s.col+d[1]
			if r < 0 || r >= m || c < 0 || c >= n || visited[r][c] {
				continue
			}
			visited[r][c] = true
			water += max(0, s.height-heightMap[r][c])
			q.Push(cell{r, c, max(s.height, heightMap[r][c])})
		}
	}
	return water
}
