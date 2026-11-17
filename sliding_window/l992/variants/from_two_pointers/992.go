package l992

// https://leetcode.com/problems/subarrays-with-k-different-integers/
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

func subarraysWithKDistinct(nums []int, k int) int {
	atMost := func(limit int) int {
		if limit < 0 {
			return 0
		}
		count := map[int]int{}
		return windowAtMost(len(nums),
			func(i int) { count[nums[i]]++ },
			func(i int) {
				count[nums[i]]--
				if count[nums[i]] == 0 {
					delete(count, nums[i])
				}
			},
			func() bool { return len(count) > limit })
	}
	return atMost(k) - atMost(k-1)
}
