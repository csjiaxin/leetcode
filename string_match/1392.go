package main

func buildPrefixArr(s string) []int {
	n := len(s)
	r := make([]int, n)
	j := 0
	for i := 1; i < n; {
		if s[i] == s[j] {
			r[i] = j + 1
			i += 1
			j += 1
		} else {
			if j == 0 {
				r[i] = 0
				i += 1
			} else {
				j = r[j-1]
			}
		}
	}
	return r
}
func longestPrefix(s string) string {
	p := buildPrefixArr(s)
	ml := p[len(s)-1]
	return s[0 : ml+1]
}
