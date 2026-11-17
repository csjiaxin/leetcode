package l1004

// https://leetcode.com/problems/max-consecutive-ones-iii/
// The standalone helper below matches the named subtemplate in README.md.
// windowLongest maintains a valid suffix after removing invalid left endpoints.
func windowLongest(n int, add, remove func(int), invalid func() bool) int {
	left, best := 0, 0
	for right := 0; right < n; right++ {
		add(right)
		for left <= right && invalid() {
			remove(left)
			left++
		}
		if length := right - left + 1; length > best {
			best = length
		}
	}
	return best
}

func longestOnes(nums []int, k int) int {
	zeros := 0
	return windowLongest(len(nums),
		func(i int) {
			if nums[i] == 0 {
				zeros++
			}
		},
		func(i int) {
			if nums[i] == 0 {
				zeros--
			}
		},
		func() bool { return zeros > k })
}
