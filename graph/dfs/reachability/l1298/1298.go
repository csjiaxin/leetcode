package main

import "fmt"

// https://leetcode.com/problems/maximum-candies-you-can-get-from-boxes/
func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	have := make([]bool, len(status))
	open := make([]bool, len(status))
	for box, value := range status {
		open[box] = value == 1
	}
	for _, box := range initialBoxes {
		have[box] = true
	}
	total := 0
	ReadyDFS(initialBoxes, func(box int) bool { return have[box] && open[box] }, func(box int) []int {
		next := make([]int, 0, len(keys[box])+len(containedBoxes[box]))
		next = append(next, keys[box]...)
		return append(next, containedBoxes[box]...)
	}, func(box int) {
		total += candies[box]
		for _, key := range keys[box] {
			open[key] = true
		}
		for _, child := range containedBoxes[box] {
			have[child] = true
		}
	})
	return total
}

// ReadyDFS 只在 canEnter 成立时标记；未就绪候选可由后续事件再次尝试。
func ReadyDFS[S comparable](starts []S, canEnter func(S) bool, next func(S) []S, visit func(S)) map[S]bool {
	visited := make(map[S]bool)
	var walk func(S)
	walk = func(state S) {
		if visited[state] || !canEnter(state) {
			return
		}
		visited[state] = true
		if visit != nil {
			visit(state)
		}
		for _, successor := range next(state) {
			walk(successor)
		}
	}
	for _, start := range starts {
		walk(start)
	}
	return visited
}

func main() {
	fmt.Println(maxCandies([]int{1, 0, 1, 0}, []int{7, 5, 4, 100}, [][]int{{}, {}, {1}, {}}, [][]int{{1, 2}, {3}, {}, {}}, []int{0}))
}
