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
func strStr(haystack string, needle string) []int {
	ret := make([]int, 0)
	prefix := buildPrefixArr(needle)
	j := 0
	i := 0
	for i < len(haystack) {
		if haystack[i] == needle[j] {

			if j == len(needle)-1 {
				//found
				// fmt.Println(i - i1)
				// return i - j
				ret = append(ret, i+1-len(needle))
				if j > 0 {
					j = prefix[j-1]
				} else {
					i += 1
				}
			} else {
				j += 1
				i += 1
			}
		} else if j > 0 {
			j = prefix[j-1]
		} else {
			i += 1
		}
	}
	return ret
}

func main() {
	s := "aaaabaa"
	fmt.Println(strStr(s, "a"))   //4
	fmt.Println(strStr(s, "aa"))  //4
	fmt.Println(strStr(s, "aaa")) //4
}
