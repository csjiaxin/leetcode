package main

import "fmt"

type Color int

const (
	White Color = iota
	Gray
	Black
)

var adj [][]int
var visit []Color
var ret bool
var retArr []int

func dfs(at int) {
	if visit[at] == White {
		visit[at] = Gray
		for _, to := range adj[at] {
			dfs(to)
		}
		visit[at] = Black
		retArr = append(retArr, at)

	} else if visit[at] == Gray {
		ret = false
	}
}
func edge(at, to int) {
	for i := 0; i < len(adj[at]); i++ {
		if adj[at][i] == to {
			return
		}
	}
	adj[at] = append(adj[at], to)
}
func sortItems(n int, m int, group []int, beforeItems [][]int) []int {

	gArr := make([][]int, m)
	for i := 0; i < n; i++ {
		if group[i] != -1 {
			gArr[group[i]] = append(gArr[group[i]], i)
		}
	}

	adj = make([][]int, n)
	ret = true
	retArr = make([]int, 0)

	for at := 0; at < n; at++ {
		for _, e := range beforeItems {
			for _, b := range e {
				edge(at, b)
				if group[b] >= 0 && group[at] >= 0 {

					for _, m := range gArr[group[b]] {
						for _, m_at := range gArr[group[at]] {
							edge(m_at, m)
						}
					}
				}
			}
		}

	}

	visit = make([]Color, n)
	for to := 0; to < n; to++ {
		dfs(to)
	}
	if ret == false {
		return []int{}
	}
	return retArr

}

func main() {
	fmt.Println(sortItems(8, 2, []int{-1, -1, 1, 0, 0, 1, 0, -1}, [][]int{{}, {6}, {5}, {6}, {3, 6}, {}, {}, {}}))
}
