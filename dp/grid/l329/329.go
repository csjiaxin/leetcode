package l329

func longestIncreasingPath(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	dir4 := [][]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	var isValid func(i, j int) bool
	isValid = func(i, j int) bool {
		if i >= 0 && i < m && j >= 0 && j < n {
			return true
		}
		return false
	}
	indegree := make([][]int, m)
	for i := 0; i < m; i++ {
		indegree[i] = make([]int, n)
	}
	q := make([][]int, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for _, d := range dir4 {
				ni := i + d[0]
				nj := j + d[1]
				if isValid(ni, nj) {
					if matrix[ni][nj] < matrix[i][j] {
						indegree[i][j] += 1
					}
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if indegree[i][j] == 0 {
				q = append(q, []int{i, j})
			}
		}
	}
	layer := 0
	for len(q) > 0 {
		L := len(q)
		layer += 1
		for i := 0; i < L; i++ {
			e := q[i]
			for _, d := range dir4 {
				ni := e[0] + d[0]
				nj := e[1] + d[1]
				if isValid(ni, nj) {
					if matrix[ni][nj] > matrix[e[0]][e[1]] {
						indegree[ni][nj] -= 1
						if indegree[ni][nj] == 0 {
							q = append(q, []int{ni, nj})
						}
					}
				}

			}
		}
		q = q[L:]
	}
	return layer
}
