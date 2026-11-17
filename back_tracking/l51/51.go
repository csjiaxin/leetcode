package l51

// https://leetcode.com/problems/n-queens/
// solveNQueens returns all valid n-queens boards in any order.
func solveNQueens(n int) [][]string {
    ret := [][]string{}
    positions := make([]int, n)

	//表示当前递归路径中，这一列是否已经有皇后。
    cols := make([]bool, n)
    diag1 := make([]bool, 2*n-1)
    diag2 := make([]bool, 2*n-1)

    var dfs func(row int)
    dfs = func(row int) {
        if row == n {
            // 所有行都放好了，生成一个独立的棋盘。
            board := make([]string, n)
            for r, c := range positions {
                line := make([]byte, n)
                for i := range line {
                    line[i] = '.'
                }
                line[c] = 'Q'
                board[r] = string(line)
            }

            ret = append(ret, board)
            return
        }

        for col := 0; col < n; col++ {
            d1 := row - col + n - 1
            d2 := row + col

            if cols[col] || diag1[d1] || diag2[d2] {
                continue
            }

            // 选择。
            positions[row] = col
            cols[col] = true
            diag1[d1] = true
            diag2[d2] = true

            // 探索。
            dfs(row + 1)

            // 撤销选择。
            cols[col] = false
            diag1[d1] = false
            diag2[d2] = false
        }
    }

    dfs(0)
    return ret
}

}
