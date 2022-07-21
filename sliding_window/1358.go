package main

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func meet(nums map[byte]int) bool {
	return nums['a'] > 0 && nums['b'] > 0 && nums['c'] > 0
}
func numberOfSubstrings(s string) int {
	nums := map[byte]int{}
	n := len(s)
	l := 0
	ret := 0
	for r := 0; r < n; r++ {
		nums[s[r]] += 1
		for ; meet(nums); l += 1 {
			ret += n - r
			nums[s[l]] -= 1
		}
	}
	return ret
}

func main() {
	println(numberOfSubstrings("abcabc"))
}
