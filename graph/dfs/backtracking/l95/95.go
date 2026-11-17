package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	r := make([]*TreeNode, 0)

	for i := 1; i <= n; i++ {
		root := dfs(i, 1, i-1, i+1, n)
		r = append(r, root...)
	}
	return r
}
func dfs(n, ll, lm, rl, rm int) []*TreeNode {

	larr := make([]*TreeNode, 0)
	rarr := make([]*TreeNode, 0)
	if ll <= lm {
		for i := ll; i <= lm; i++ {
			larr = append(larr, dfs(i, ll, i-1, i+1, lm)...)
		}
	} else {
		larr = append(larr, nil)
	}
	if rl <= rm {
		for i := rl; i <= rm; i++ {
			rarr = append(rarr, dfs(i, rl, i-1, i+1, rm)...)
		}

	} else {
		rarr = append(rarr, nil)
	}
	retarr := make([]*TreeNode, 0)
	for _, l := range larr {
		for _, r := range rarr {
			node := &TreeNode{n, l, r}
			retarr = append(retarr, node)
		}
	}
	return retarr
}

func main() {
	fmt.Println(generateTrees(1))
	fmt.Println(generateTrees(2))
	fmt.Println(generateTrees(3))
}
