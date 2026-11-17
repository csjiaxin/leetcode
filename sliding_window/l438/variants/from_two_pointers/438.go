package l438

// https://leetcode.com/problems/find-all-anagrams-in-a-string/
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

func findAnagrams(s string, p string) []int {
	result := []int{}
	if len(p) == 0 {
		return result
	}
	var need, count [256]int
	for i := range p {
		need[p[i]]++
	}
	windowFixed(len(s), len(p),
		func(i int) { count[s[i]]++ },
		func(i int) { count[s[i]]-- },
		func(left int) {
			if count == need {
				result = append(result, left)
			}
		})
	return result
}
