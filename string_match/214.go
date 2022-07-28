package main

import "fmt"

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

func reverse(s string) string {
	runes := []byte(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func shortestPalindrome(s string) string {
	tmp := s + "#" + reverse(s)
	prefix := buildPrefixArr(tmp)
	ml := prefix[len(prefix)-1]

	return reverse(s[ml:]) + s
}

func main() {
	fmt.Println(shortestPalindrome("aacecaaa"))
}
