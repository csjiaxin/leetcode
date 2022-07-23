package main

import "fmt"

var r map[int]int

func find(b int) (int, bool) {
	f, ok := r[b]
	if !ok {
		return 0, false
	}
	if f == b {
		return b, true
	}
	return find(f)
}
func union_if_exist(n, b int) {
	if p, ok := find(b); ok {
		np, _ := find(n)
		r[p] = np
		// rsize[np] += rsize[p]
	}
}
func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	r = map[int]int{}

	for i := 0; i < n; i++ {
		r[i] = i
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected[i][j] > 0 {

				union_if_exist(i, j)
			}
		}
	}
	fmap := map[int]int{}
	for i := 0; i < n; i++ {
		ip, _ := find(i)
		fmap[ip] += 1
	}
	return len(fmap)
}

func main() {
	fmt.Println(findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))
}
