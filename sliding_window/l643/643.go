package l643

// https://leetcode.com/problems/maximum-average-subarray-i/
// The standalone helper below matches the named subtemplate in README.md.
// windowFixed visits exactly the windows of length k.
func windowFixed(n, k int, add, remove func(int), visit func(int)) {
	if k <= 0 || k > n {
		return
	}
	left := 0
	for right := 0; right < n; right++ {
		add(right)
		if right-left+1 > k {
			remove(left)
			left++
		}
		if right-left+1 == k {
			visit(left)
		}
	}
}

func findMaxAverage(nums []int, k int) float64 {
	if k <= 0 || k > len(nums) {
		return 0
	}
	var sum, best int64
	first := true
	windowFixed(len(nums), k,
		func(i int) { sum += int64(nums[i]) },
		func(i int) { sum -= int64(nums[i]) },
		func(left int) {
			if first || sum > best {
				best = sum
				first = false
			}
		})
	return float64(best) / float64(k)
}
