package main
import "fmt"

func dfs(at int, status []int, keys [][]int, containedBoxes [][]int) {
	if status[at]&Visit != 0 {
		return
	}
	if status[at]|Have != 0 && status[at]|Open != 0 {

		status[at] |= Visit

		for _, c := range containedBoxes[at] {
			status[c] |= Have
			dfs(c, status, keys, containedBoxes)
		}
		for _, c := range keys[at] {
			status[c] |= Open
			dfs(c, status, keys, containedBoxes)
		}
	}
}

const (
	Open  = 1
	Have  = 2
	Visit = 4
)

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	for _, v := range initialBoxes {
		status[v] |= Have
		dfs(v, status, keys, containedBoxes)
	}
	ret := 0
	for i, v := range status {
		if v&Open != 0 && v&Have != 0 {
			ret += candies[i]
		}
	}
	return ret
}

func main() {
	fmt.Println(maxCandies([]int{1, 0, 1, 0}, []int{7, 5, 4, 100}, [][]int{{}, {}, {1}, {}}, [][]int{{1, 2}, {3}, {}, {}}, []int{0}))
}
