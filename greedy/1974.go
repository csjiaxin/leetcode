apackage main

import "fmt"

func minTimeToType(word string) int {
	point := byte('a')
	n := 0
	for i := 0; i < len(word); i++ {
		c := word[i]
		diff := int(c) - int(point)
		if diff < 0 {
			diff = -diff
		}
		if diff > 13 {
			diff = 26 - diff
		}
		n += diff + 1
		point = c
	}
	return n

}

func main() {
	fmt.Println(minTimeToType("bza"))
}
