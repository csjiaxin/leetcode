package l904

// https://leetcode.com/problems/fruit-into-baskets/
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

func totalFruit(nums []int) int {
	count := map[int]int{}
	return windowLongest(len(nums),
		func(i int) { count[nums[i]]++ },
		func(i int) {
			count[nums[i]]--
			if count[nums[i]] == 0 {
				delete(count, nums[i])
			}
		},
		func() bool { return len(count) > 2 })
}
