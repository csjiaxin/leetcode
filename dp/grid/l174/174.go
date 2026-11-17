package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

var m map[string]int
var dir = [][]int{{1, 0}, {0, 1}}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(dungeon [][]int, i, j int) int {
	if i == len(dungeon)-1 && j == len(dungeon[0])-1 {
		return 1 - dungeon[i][j]
	}
	key := fmt.Sprintf("%d,%d", i, j)
	if v, ok := m[key]; ok {
		return v
	}
	ans := 99999999
	for _, d := range dir {
		ni := i + d[0]
		nj := j + d[1]
		if ni < len(dungeon) && nj < len(dungeon[0]) {
			ans = min(ans, max(dfs(dungeon, ni, nj), 1))
		}
	}
	v := ans - dungeon[i][j]
	m[key] = v
	return v
}
func calculateMinimumHP(dungeon [][]int) int {
	m = map[string]int{}
	return max(dfs(dungeon, 0, 0), 1)
}

func main() {
	fmt.Println(calculateMinimumHP([][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}))
}
