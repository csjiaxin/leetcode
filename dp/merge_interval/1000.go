package main

import (
	"fmt"
	"math"
	"strconv"
)

var cost int

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func get_state(stones []int) string {
	s := ""
	for _, v := range stones {
		s += strconv.Itoa(v)
	}
	return s
}

var save map[string]int

func dfs(stones []int, k int) int {
	if len(stones) == 1 {
		return 0
	}
	if len(stones) < k {
		return -1
	}
	state := get_state(stones)
	if v, ok := save[state]; ok {
		return v
	}

	_min := math.MaxInt
	for i := 0; i+k <= len(stones); i++ {
		sum := 0
		for j := 0; j < k; j++ {
			sum += stones[i+j]
		}
		n := make([]int, len(stones)-k+1)
		copy(n, stones[0:i])
		// n = append(n, stones[0:i]...)
		// n = append(n, sum)
		n[i] = sum
		if i+k < len(stones) {
			// n = append(n, stones[i+k:]...)
			copy(n[i+1:], stones[i+k:])
		}
		solution := dfs(n, k)
		if solution < 0 {
			continue
		}
		_min = min(_min, sum+solution)
	}
	if _min == math.MaxInt {
		_min = -1
	}

	save[state] = _min
	return _min

}
func mergeStones(stones []int, k int) int {
	cost = math.MaxInt
	save = map[string]int{}
	return dfs(stones, k)
}

func main() {
	fmt.Println(mergeStones([]int{16, 43, 87, 30, 4, 98, 12, 30, 47, 45, 32, 4, 64, 14, 24, 84, 86, 51, 11, 22, 4}, 2))
	fmt.Println(mergeStones([]int{3, 2, 4, 1}, 3))
}
