package l1760

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

func minimumSize(nums []int, maxOperations int) int {

	r := slices.Max(nums)

	p := firstTrue(1, r, func(penalty int) bool {

		needOp := 0
		for _, n := range nums {
			needOp += (n - 1) / penalty
		}
		return needOp <= maxOperations
	})
	return p
}
