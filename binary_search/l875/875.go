package l875

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

func minEatingSpeed(piles []int, h int) int {
	r := slices.Max(piles) + 1

	p := firstTrue(1, r, func(speed int) bool {

		needOp := 0
		for _, n := range piles {
			needOp += (n + speed - 1) / speed
		}
		return needOp <= h
	})
	return p

}
