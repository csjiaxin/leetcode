package l209

// https://leetcode.com/problems/minimum-size-subarray-sum/
// The standalone helper below matches the named subtemplate in README.md.
// windowShortest records valid windows before shrinking; length n+1 means absent.
func windowShortest(n int, add, remove func(int), valid func() bool) (int, int) {
	left, start, length := 0, 0, n+1
	for right := 0; right < n; right++ {
		add(right)
		for left <= right && valid() {
			if right-left+1 < length {
				start, length = left, right-left+1
			}
			remove(left)
			left++
		}
	}
	return start, length
}

func minSubArrayLen(target int, nums []int) int {
	var sum int64
	_, length := windowShortest(len(nums),
		func(i int) { sum += int64(nums[i]) },
		func(i int) { sum -= int64(nums[i]) },
		func() bool { return sum >= int64(target) })
	if length == len(nums)+1 {
		return 0
	}
	return length
}
