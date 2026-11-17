package main

// ReverseGrid 从右下倒推；cell 只能读取已完成的下、右状态。
func ReverseGrid[T any](rows, cols int, cell func(r, c int, down, right T) T) [][]T {
	dp := make([][]T, rows)
	for r := range dp {
		dp[r] = make([]T, cols)
	}
	for r := rows - 1; r >= 0; r-- {
		for c := cols - 1; c >= 0; c-- {
			var down, right T
			if r+1 < rows {
				down = dp[r+1][c]
			}
			if c+1 < cols {
				right = dp[r][c+1]
			}
			dp[r][c] = cell(r, c, down, right)
		}
	}
	return dp
}

// 状态是进入当前格所需的最低血量，反向推导且始终至少为 1。
func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 1
	}
	m, n := len(dungeon), len(dungeon[0])
	dp := ReverseGrid(m, n, func(r, c, down, right int) int {
		if r == m-1 && c == n-1 {
			return max(1, 1-dungeon[r][c])
		}
		need := int(^uint(0) >> 1)
		if r+1 < m {
			need = min(need, down)
		}
		if c+1 < n {
			need = min(need, right)
		}
		return max(1, need-dungeon[r][c])
	})
	return dp[0][0]
}
