package l1011

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

func shipWithinDays(weights []int, days int) int {
	slices.Sort(weights)
	r := slices.Max(weights)
	N := len(weights)

	p := firstTrue(1, r, func(capacity int) bool {

		needOp := 0

		for j := N - 1; j > 0; j-- {

		}

		return needOp <= days
	})
	return p
}
