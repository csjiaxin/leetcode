package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

var memo map[string]int

func dfs(coins []int, l int, start int) int {
	if l == 0 {
		return 1
	}
	key := fmt.Sprintf("%d,%d", l, start)
	if v, ok := memo[key]; ok {
		return v
	}
	ans := 0
	for i := start; i < len(coins); i++ {
		v := coins[i]
		if v <= l {
			ans += dfs(coins, l-v, i)
		}
	}
	memo[key] = ans
	return ans
}
func change(amount int, coins []int) int {
	memo = map[string]int{}
	return dfs(coins, amount, 0)
}

func main() {
	fmt.Println(change(5, []int{1, 2, 5})) //4
}
