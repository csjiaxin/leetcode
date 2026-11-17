package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minimumTime(n int, relations [][]int, time []int) int {
	N := n + 1
	from := make([][]int, N)
	to := make([][]int, N)
	indegre := make([]int, N)
	for _, r := range relations {
		prev := r[0]
		next := r[1]
		indegre[next] += 1
		from[next] = append(from[next], prev)
		to[prev] = append(to[prev], next)
	}
	q := []int{}
	dp := make([]int, N)
	for i := 1; i < N; i++ {
		if indegre[i] == 0 {
			q = append(q, i)
		}
	}
	ret := 0
	for len(q) > 0 {
		s := len(q)
		for i := 0; i < s; i++ {
			front := q[0]
			q = q[1:]
			dp[front] = time[front-1]
			ret = max(dp[front], ret)
			for _, prev := range from[front] {
				dp[front] = max(dp[front], dp[prev]+time[front-1])
				ret = max(dp[front], ret)
			}

			for _, next := range to[front] {
				indegre[next] -= 1
				if indegre[next] == 0 {
					q = append(q, next)
				}
			}

		}
	}
	return ret
}

func main() {
	// fmt.Println(minimumTime(3, [][]int{{1, 3}, {2, 3}}, []int{3, 2, 5}))
	fmt.Println(minimumTime(1, [][]int{}, []int{1}))
}
