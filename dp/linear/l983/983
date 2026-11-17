package main

import "fmt"

var memo map[string]int

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func dfs(days, costs []int, i int, end int) int {
	if i >= len(days) {
		return 0
	}
	key := fmt.Sprintf("%d,%d", i, end)
	if v, ok := memo[key]; ok {
		return v
	}
	ans := 99999999
	if days[i] <= end {
		ans = dfs(days, costs, i+1, end)
	} else {
		for j := 0; j < len(costs); j++ {
			l := 0
			switch j {
			case 0:
				l = 1
			case 1:
				l = 7
			case 2:
				l = 30
			}
			ans = min(ans, dfs(days, costs, i+1, days[i]+l-1)+costs[j])
		}
	}
	memo[key] = ans
	return ans
}
func mincostTickets(days []int, costs []int) int {
	memo = map[string]int{}
	return dfs(days, costs, 0, 0)
}
func main() {
	fmt.Println(mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15}))
}
