package main

import "fmt"

func compare(m1, m2 map[byte]int) int {
	for k, v := range m1 {
		if m2[k] < v {
			return -1
		} else if m2[k] > v {
			return 1
		}
	}
	return 0
}
func checkInclusion(s1 string, s2 string) bool {

	l := 0
	m1 := map[byte]int{}
	m2 := map[byte]int{}
	for i := 0; i < len(s1); i++ {
		m1[s1[i]] += 1
	}
	for r := 0; r < len(s2); r++ {
		c := s2[r]
		if _, ok := m1[c]; !ok {
			l = r + 1
			m2 = map[byte]int{}
			continue
		}
		m2[c] += 1
		for compare(m1, m2) > 0 {
			m2[s2[l]] -= 1
			l += 1
		}
		if compare(m1, m2) == 0 {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(checkInclusion("adc", "dcda"))
}
