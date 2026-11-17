package l827

import "leetcode/graph/dfs/grid"

// FloodFill 使用共享访问表处理一个四连通块，不修改原网格。
func FloodFill(rows, cols, startR, startC int, visited [][]bool, canEnter func(r, c int) bool, visit func(r, c int)) {
	grid.FloodFill(rows, cols, startR, startC, visited, canEnter, visit)
}

// https://leetcode.com/problems/making-a-large-island/
func largestIsland(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	labels := make([][]int, rows)
	for r := range visited {
		visited[r] = make([]bool, cols)
		labels[r] = make([]int, cols)
	}
	areas := map[int]int{}
	id, best := 0, 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != 1 || visited[r][c] {
				continue
			}
			id++
			FloodFill(rows, cols, r, c, visited,
				func(r, c int) bool { return grid[r][c] == 1 },
				func(r, c int) {
					labels[r][c] = id
					areas[id]++
				},
			)
			if areas[id] > best {
				best = areas[id]
			}
		}
	}
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] != 0 {
				continue
			}
			area := 1
			seen := map[int]bool{}
			for _, d := range directions {
				nr, nc := r+d[0], c+d[1]
				if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
					continue
				}
				neighborID := labels[nr][nc]
				// 按编号去重：同一块不能重复，不同块面积相同也都要加。
				if neighborID != 0 && !seen[neighborID] {
					seen[neighborID] = true
					area += areas[neighborID]
				}
			}
			if area > best {
				best = area
			}
		}
	}
	return best
}
