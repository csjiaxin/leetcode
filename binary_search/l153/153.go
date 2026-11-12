package l153

func findMin(nums []int) int {
	N := len(nums)
	target := nums[N-1]
	lessEqual := func(x int) bool {
		return x <= target
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

	l := bs(lessEqual)
	return nums[l]
}
