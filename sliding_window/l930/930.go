package l930

// https://leetcode.com/problems/binary-subarrays-with-sum/
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

func numSubarraysWithSum(nums []int, goal int) int {
	atMost := func(limit int) int {
		if limit < 0 {
			return 0
		}
		sum := 0
		return windowAtMost(len(nums),
			func(i int) { sum += nums[i] },
			func(i int) { sum -= nums[i] },
			func() bool { return sum > limit })
	}
	return atMost(goal) - atMost(goal-1)
}
