package l567

// https://leetcode.com/problems/permutation-in-string/
// The standalone helper below matches the named subtemplate in README.md.
// windowFixed visits exactly the windows of length k.
func windowFixed(n, k int, add, remove func(int), visit func(int)) {
	if k <= 0 || k > n {
		return
	}
	left := 0
	for right := 0; right < n; right++ {
		add(right)
		if right-left+1 > k {
			remove(left)
			left++
		}
		if right-left+1 == k {
			visit(left)
		}
	}
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) == 0 {
		return true
	}
	var need, count [256]int
	for i := range s1 {
		need[s1[i]]++
	}
	found := false
	windowFixed(len(s2), len(s1),
		func(i int) { count[s2[i]]++ },
		func(i int) { count[s2[i]]-- },
		func(left int) {
			if count == need {
				found = true
			}
		})
	return found
}
