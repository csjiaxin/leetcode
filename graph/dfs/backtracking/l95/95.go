package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// https://leetcode.com/problems/unique-binary-search-trees-ii/
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	solve := MemoDFS(func(interval [2]int, recur func([2]int) []*TreeNode) []*TreeNode {
		left, right := interval[0], interval[1]
		if left > right {
			return []*TreeNode{nil}
		}
		result := make([]*TreeNode, 0)
		for root := left; root <= right; root++ {
			for _, l := range recur([2]int{left, root - 1}) {
				for _, r := range recur([2]int{root + 1, right}) {
					result = append(result, &TreeNode{Val: root, Left: l, Right: r})
				}
			}
		}
		return result
	})
	return solve([2]int{1, n})
}

func MemoDFS[S comparable, V any](evaluate func(S, func(S) V) V) func(S) V {
	memo := make(map[S]V)
	var solve func(S) V
	solve = func(state S) V {
		if value, ok := memo[state]; ok {
			return value
		}
		value := evaluate(state, solve)
		memo[state] = value
		return value
	}
	return solve
}

func main() { fmt.Println(len(generateTrees(1)), len(generateTrees(2)), len(generateTrees(3))) }
