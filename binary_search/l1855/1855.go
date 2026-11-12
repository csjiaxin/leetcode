package l1855

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

func maxDistance(nums1, nums2 []int) int {
	N2 := len(nums2)
	ret := 0
	for i, n1 := range nums1 {

		p2 := firstTrue(i, N2-1, func(i int) bool {
			return nums2[i] < n1
		}) - 1

		dis := p2 - i
		ret = max(ret, dis)
	}
	return ret
}
