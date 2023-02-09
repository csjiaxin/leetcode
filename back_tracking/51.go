package main

import "strings"
import "fmt"

var res [][]string

func solveNQueens(n int) [][]string {
	res = make([][]string, 0)
	board := make([]int, 0)
	dfs(0, board, n)
	return res
}

func drawBoard(board []int) []string {
	r := []string{}
	for _, v := range board {
		_s := strings.Repeat(".", v) + "Q" + strings.Repeat(".", len(board)-1-v)
		r = append(r, _s)
	}
	return r
}
func not_in(a int, arr []int) bool {
	for _, v := range arr {
		if v == a {
			return false
		}
	}
	return true
}
func not_dia(a int, arr []int)bool{
	for k
}
func dfs(i int, board []int, n int) {
	if i == n {
		res = append(res, drawBoard(board))
		return
	}
	for j := 0; j < n; j++ {
		if not_in(j, board) && not_dia(j, board) {
			board = append(board, j)
			dfs(i+1, board, n)
			board = board[:len(board)-1]
		}
	}

}

func main() {

	fmt.Println(solveNQueens(4))
}
