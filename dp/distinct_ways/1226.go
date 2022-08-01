package main

import (
	"fmt"
	"math"
)

var m map[string]int

func dfs(l, arrLen, at int) int {
	if l == 0 && at == 0 {
		return 1
	}
	if l == 0 {
		return 0
	}
	key := fmt.Sprintf("%d,%d", l, at)
	if v, ok := m[key]; ok {
		return v
	}
	ans := 0
	if at > 0 {

		ans += dfs(l-1, arrLen, at-1)
		ans = int(math.Mod(float64(ans), 1e9+7))
	}
	if at+1 < arrLen {
		ans += dfs(l-1, arrLen, at+1)
		ans = int(math.Mod(float64(ans), 1e9+7))
	}
	ans += dfs(l-1, arrLen, at)
	ans = int(math.Mod(float64(ans), 1e9+7))
	m[key] = ans
	return ans

}
func numWays(steps int, arrLen int) int {
	m = map[string]int{}
	ans := 0
	// for i := 0; i < arrLen; i++ {
	ans += dfs(steps, arrLen, 0)
	ans = int(math.Mod(float64(ans), 1e9+7))
	// }
	return ans
}
func main() {
	fmt.Println(numWays(3, 2))
}
