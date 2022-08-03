package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

var memo map[string]int

func dfs(coins []int, l int) int {
	if l == 0 {
		return 0
	}
	key := fmt.Sprintf("%d", l)
	if v, ok := memo[key]; ok {
		return v
	}
	ans := -1
	for _, v := range coins {
		if v <= l {
			step := dfs(coins, l-v) + 1
			if step == 0 {
				continue
			}
			if ans < 0 {
				ans = step
			} else {
				ans = min(ans, step)
			}
		}
	}
	memo[key] = ans
	return ans
}
func coinChange(coins []int, amount int) int {
	memo = map[string]int{}
	return dfs(coins, amount)
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}
