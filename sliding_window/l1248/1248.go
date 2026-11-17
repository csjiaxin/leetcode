package l1248

// https://leetcode.com/problems/count-number-of-nice-subarrays/
// The standalone helper below matches the named subtemplate in README.md.
// windowAtMost counts valid suffixes; removing the left endpoint preserves validity.
func windowAtMost(n int, add, remove func(int), invalid func() bool) int {
	left, count := 0, 0
	for right := 0; right < n; right++ {
		add(right)
		for left <= right && invalid() {
			remove(left)
			left++
		}
		count += right - left + 1
	}
	return count
}

func numberOfSubarrays(nums []int, k int) int {
	atMost := func(limit int) int {
		if limit < 0 {
			return 0
		}
		odd := 0
		return windowAtMost(len(nums),
			func(i int) { odd += nums[i] & 1 },
			func(i int) { odd -= nums[i] & 1 },
			func() bool { return odd > limit })
	}
	return atMost(k) - atMost(k-1)
}
