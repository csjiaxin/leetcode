package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func get_max(a []int) int {
	_max := a[0]
	for i := 0; i < len(a); i++ {
		_max = max(_max, a[i])
	}
	return _max
}

var ret int

func dfs(i int, cookies []int, kpack []int) {
	if i >= len(cookies) {
		_max := get_max(kpack)
		ret = min(_max, ret)
		return
	}
	for j := 0; j < len(kpack); j++ {
		kpack[j] += cookies[i]
		dfs(i+1, cookies, kpack)
		kpack[j] -= cookies[i]
	}
}
func distributeCookies(cookies []int, k int) int {
	ret = math.MaxInt
	kpack := make([]int, k)

	dfs(0, cookies, kpack)
	return ret
}

func main() {
	fmt.Println(distributeCookies([]int{8, 15, 10, 20, 8}, 2))
}
