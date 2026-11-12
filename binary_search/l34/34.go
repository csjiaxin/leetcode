package l34

func searchRange(nums []int, target int) []int {

	firstGreaterEqual := func(x int) bool {
		return x >= target
	}
	firstGreater := func(x int) bool {
		return x > target
	}

	bs := func(check func(x int) bool) int {
		l, r := 0, len(nums)-1

		for l <= r {
			mid := l + (r-l)/2

			if check(nums[mid]) {
				r = mid - 1
			} else {
				l = mid + 1
			}

		}
		return l
	}
	firstGreaterEqualRet := bs(firstGreaterEqual)
	firstGreaterRet := bs(firstGreater)
	N := len(nums)
	if firstGreaterEqualRet <= N-1 {
		if nums[firstGreaterEqualRet] == target {
			return []int{firstGreaterEqualRet, firstGreaterRet - 1}
		}
	}
	return []int{-1, -1}
}
