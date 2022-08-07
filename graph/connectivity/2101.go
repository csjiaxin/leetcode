package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func dfs(i int, visit []bool, to [][]int) {

	for _, v := range to[i] {
		if visit[v] == false {

			visit[v] = true
			dfs(v, visit, to)
		}
	}
}

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	from := make([][]int, n)
	to := make([][]int, n)
	for i := 0; i < len(bombs); i++ {
		for j := 0; j < len(bombs); j++ {
			bi := bombs[i]
			bj := bombs[j]
			x := bi[0] - bj[0]
			y := bi[1] - bj[1]
			d2 := (x * x) + y*y
			// i->j
			if bi[2]*bi[2] >= d2 {
				from[j] = append(from[j], i)
				to[i] = append(to[i], j)
			}
			// j->i
			if bj[2]*bj[2] >= d2 {
				from[i] = append(from[i], j)
				to[j] = append(to[j], i)
			}
		}
	}
	ret := 0
	for i := 0; i < n; i++ {
		visit := make([]bool, n)
		visit[i] = true
		dfs(i, visit, to)
		count := 0
		for _, v := range visit {
			if v {
				count += 1
			}
		}
		ret = max(ret, count)

	}
	return ret
}

func main() {
	fmt.Println(maximumDetonation([][]int{{2, 1, 3}, {6, 1, 4}}))
}
