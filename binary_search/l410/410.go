package l410

import "slices"

func firstTrue(l, r int, check func(int) bool) int {

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
func splitArray(nums []int, k int) int {
	total := 0
	for _, n := range nums {
		total += n
	}

	p := firstTrue(slices.Max(nums), total, func(largest_sum int) bool {
		need := 1
		load := 0
		for _, n := range nums {
			if load+n > largest_sum {
				need++
				load = 0
			}
			load += n
		}
		return need <= k
	})
	return p
}
