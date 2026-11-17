package l713

// https://leetcode.com/problems/subarray-product-less-than-k/
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

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	product := int64(1)
	return windowAtMost(len(nums),
		func(i int) { product *= int64(nums[i]) },
		func(i int) { product /= int64(nums[i]) },
		func() bool { return product >= int64(k) })
}
