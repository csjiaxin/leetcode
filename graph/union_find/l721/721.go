package l721

import "sort"

// https://leetcode.com/problems/accounts-merge/
func accountsMerge(accounts [][]string) [][]string {
	uf := NewUnionFind(len(accounts))
	owner := make(map[string]int)
	for i, account := range accounts {
		for _, email := range account[1:] {
			if previous, ok := owner[email]; ok {
				uf.Union(i, previous)
			} else {
				owner[email] = i
			}
		}
	}
	groups := make(map[int][]string)
	for email, account := range owner {
		root := uf.Find(account)
		groups[root] = append(groups[root], email)
	}
	answer := make([][]string, 0, len(groups))
	for root, emails := range groups {
		sort.Strings(emails)
		group := append([]string{accounts[root][0]}, emails...)
		answer = append(answer, group)
	}
	return answer
}

type UnionFind struct {
	parent []int
	size   []int
}

// Nodes are numbered 0 through n-1.
func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		size:   make([]int, n),
	}

	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}

	return uf
}

func (uf *UnionFind) Find(x int) int {
	for x != uf.parent[x] {
		// Path compression: point x directly to its grandparent.
		uf.parent[x] = uf.parent[uf.parent[x]]
		x = uf.parent[x]
	}
	return x
}

func (uf *UnionFind) Union(a, b int) bool {
	rootA := uf.Find(a)
	rootB := uf.Find(b)

	if rootA == rootB {
		return false
	}

	// Attach the smaller tree under the larger tree.
	if uf.size[rootA] < uf.size[rootB] {
		rootA, rootB = rootB, rootA
	}

	uf.parent[rootB] = rootA
	uf.size[rootA] += uf.size[rootB]
	return true
}

func (uf *UnionFind) Connected(a, b int) bool {
	return uf.Find(a) == uf.Find(b)
}
