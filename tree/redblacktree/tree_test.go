package redblacktree

import (
	"github.com/emirpasic/gods/trees/redblacktree"
	"testing"
)

func TestTree(t *testing.T) {

	arr := []int{1, 4, 6, 3, 5, 7, 8, 2, 9}
	tree := NewWithIntComparator()
	tree2 := redblacktree.NewWithIntComparator()
	for _, i := range arr {
		// save := tree.StringWithColor()
		// t.Logf("Before: %v", save)
		tree.Put(i, i)
		// t.Logf("After: %v", tree.StringWithColor())
		tree2.Put(i, i)
		if tree.String() != tree2.String() {
			t.Fatal(tree.StringWithColor(), tree2.String())
		}
		if tree.Left().Value != tree2.Left().Value {
			t.Fatal("Left:", tree.Left().Value, tree2.Left().Value)

		}
		if tree.Right().Value != tree2.Right().Value {
			t.Fatal("Right:")
		}
	}
	// print(tree.String())
	// t.Error(tree.String())
}
