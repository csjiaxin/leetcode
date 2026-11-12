package l1608

import "sort"

func firstTrue(n int, check func(int) bool) int {
	l, r := 0, n-1

	for l <= r {
		mid := l + (r-l)/2

		if check(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l
}

func specialArray(nums []int) int {
	sort.Ints(nums)
	n := len(nums)

	count := func(x int) int {
		p := firstTrue(n, func(i int) bool {
			return nums[i] >= x
		})
		return n - 1 - p + 1
	}

	candidate := firstTrue(n+1, func(x int) bool {
		return count(x) <= x
	})
	if count(candidate) == candidate {
		return candidate
	}

	return -1

}
