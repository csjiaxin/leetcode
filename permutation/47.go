package main

import "fmt"

var result [][]int

func dfs(nums, path []int, at int, visit []bool) {
	if at >= 0 {

		visit[at] = true
		path = append(path, nums[at])
	}
	if len(path) >= len(nums) {
		c := make([]int, len(path))
		copy(c, path)
		result = append(result, c)
	} else {
		choosen := map[int]bool{}
		for i := 0; i < len(nums); i++ {
			if !visit[i] && choosen[nums[i]] == false {
				choosen[nums[i]] = true
				dfs(nums, path, i, visit)
			}
		}
	}
	if at >= 0 {

		visit[at] = false
	}
	// path = path[:len(path)-1]

}

func permuteUnique(nums []int) [][]int {
	result = make([][]int, 0)
	visit := make([]bool, len(nums))
	dfs(nums, []int{}, -1, visit)
	return result
}
func main() {
	fmt.Println(permuteUnique([]int{1, 2, 1}))
}
