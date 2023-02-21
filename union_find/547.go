package main

import "fmt"

var r map[int]int

func find(b int) int {
	f := r[b]
	if f == b {
		return b
	}
	r[b] = find(f)
	return r[b]
}
func union_if_exist(n, b int) {
	bp := find(b)
	np := find(n)
	r[bp] = np
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
		ip := find(i)
		fmap[ip] += 1
	}
	return len(fmap)
}

func main() {
	fmt.Println(findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))
}
