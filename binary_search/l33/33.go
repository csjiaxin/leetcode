package l33

func search(nums []int, target int) int {

	N := len(nums)
	if N == 0 {
		return -1
	}
	last := nums[N-1]

	inbetween := func(x int) bool {
		return x >= target && x <= last
	}
	notInBetween := func(x int) bool {
		return x <= last || x >= target
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
	idx := N - 1
	if target < last {
		idx = bs(inbetween)
	} else if target > last {
		idx = bs(notInBetween)
	}
	if idx >= N {
		return -1
	}
	if nums[idx] == target {
		return idx
	}
	return -1
}
