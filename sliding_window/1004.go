package main

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func longestOnes(nums []int, k int) int {

	l := 0
	ret := 0
	n := len(nums)
	r := 0
	for ; r < n; r++ {
		if nums[r] == 0 {
			ret = max(ret, r-l)
			for ; k == 0 && l <= r; l += 1 {
				if nums[l] == 0 {
					k += 1
				}
			}
			k -= 1
		}
	}

	ret = max(ret, r-l)

	return ret
}

func main() {
	println(longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	println(longestOnes([]int{0, 0, 0, 1}, 4))
}
