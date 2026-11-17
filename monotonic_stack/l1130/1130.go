package l1130

// https://leetcode.com/problems/minimum-cost-tree-from-leaf-values/
// Embedded subtemplate: see README.md for its invariant and tie policy.
func minimumLeafTreeCost(arr []int) int {
	stack := []int{}
	cost := 0
	for _, value := range arr {
		for len(stack) > 0 && stack[len(stack)-1] <= value {
			middle := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			neighbor := value
			if len(stack) > 0 && stack[len(stack)-1] < neighbor {
				neighbor = stack[len(stack)-1]
			}
			cost += middle * neighbor
		}
		stack = append(stack, value)
	}
	for len(stack) > 1 {
		last := len(stack) - 1
		cost += stack[last] * stack[last-1]
		stack = stack[:last]
	}
	return cost
}

func mctFromLeafValues(arr []int) int { return minimumLeafTreeCost(arr) }
