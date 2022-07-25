// References: http://en.wikipedia.org/wiki/Red%E2%80%93black_tree
// https://www.youtube.com/playlist?list=PLpPXw4zFa0uKKhaSz87IowJnOTzh9tiBk
// https://www.geeksforgeeks.org/avl-tree-set-1-insertion/
package avltree

import (
	"fmt"

	"github.com/emirpasic/gods/utils"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type Tree struct {
	Root       *Node
	size       int
	Comparator utils.Comparator
}

// Node is a single element within the tree
type Node struct {
	Key   interface{}
	Value interface{}
	Left  *Node
	Right *Node
	// Parent *Node
	height int
}

// NewWith instantiates a red-black tree with the custom comparator.
func NewWith(comparator utils.Comparator) *Tree {
	return &Tree{Comparator: comparator}
}

// NewWithIntComparator instantiates a red-black tree with the IntComparator, i.e. keys are of type int.
func NewWithIntComparator() *Tree {
	return &Tree{Comparator: utils.IntComparator}
}

// NewWithStringComparator instantiates a red-black tree with the StringComparator, i.e. keys are of type string.
func NewWithStringComparator() *Tree {
	return &Tree{Comparator: utils.StringComparator}
}

func (tree *Tree) insert(p *Node, n *Node) *Node {
	fmt.Println("insert:", p, n)
	if p == nil {
		return n
	}
	tree.size = tree.size + 1
	c := tree.Comparator(n.Key, p.Key)
	if c < 0 {
		p.Left = tree.insert(p.Left, n)
	} else if c > 0 {
		p.Right = tree.insert(p.Right, n)
	} else {
		p.Value = n.Value
	}
	p.updateHeight()
	return tree.checkBalance(p)
}

// Put inserts node into the tree.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree) Put(key interface{}, value interface{}) {
	n := &Node{Key: key, Value: value, height: 1}
	tree.Root = tree.insert(tree.Root, n)
}
func (tree *Tree) checkBalance(n *Node) *Node {
	if n == nil {
		return n
	}
	d := n.Left.Height() - n.Right.Height()
	fmt.Println("checkBalance:", n, d)
	if -1 <= d && d <= 1 {
		return n
	}
	if d < -1 {
		if n.Right.Left.Height() < n.Right.Right.Height() {
		} else {
			n.Right = tree.rotateRight(n.Right)
		}
		return tree.rotateLeft(n)
	} else if d > 1 {
		if n.Left.Left.Height() < n.Left.Right.Height() {
			n.Left = tree.rotateLeft(n.Left)
		}
		return tree.rotateRight(n)
	}
	return n
}

func (tree *Tree) rotateLeft(node *Node) *Node {
	child := node.Right
	// tree.updateNodeParent(node, child)
	node.Right = child.Left
	// if child.Left != nil {
	// 	child.Left.Parent = node
	// }
	child.Left = node
	// node.Parent = child
	node.updateHeight()
	child.updateHeight()
	return child
}

func (tree *Tree) rotateRight(node *Node) *Node {
	child := node.Left
	// tree.updateNodeParent(node, child)
	node.Left = child.Right
	// if child.Right != nil {
	// 	child.Right.Parent = node
	// }
	child.Right = node
	// node.Parent = child
	node.updateHeight()
	child.updateHeight()
	return child
}

// Get searches the node in the tree by key and returns its value or nil if key is not found in tree.
// Second return parameter is true if key was found, otherwise false.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree) Get(key interface{}) (value interface{}, found bool) {
	node := tree.lookup(key)
	if node != nil {
		return node.Value, true
	}
	return nil, false
}

// GetNode searches the node in the tree by key and returns its node or nil if key is not found in tree.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree) GetNode(key interface{}) *Node {
	return tree.lookup(key)
}

// Size returns the number of elements stored in the subtree.
// Computed dynamically on each call, i.e. the subtree is traversed to count the number of the nodes.
func (node *Node) Size() int {
	if node == nil {
		return 0
	}
	size := 1
	if node.Left != nil {
		size += node.Left.Size()
	}
	if node.Right != nil {
		size += node.Right.Size()
	}
	return size
}

// Left returns the left-most (min) node or nil if tree is empty.
func (tree *Tree) Left() *Node {
	var parent *Node
	current := tree.Root

	for current != nil {
		parent = current
		current = current.Left
	}
	return parent

}

// Right returns the right-most (max) node or nil if tree is empty.
func (tree *Tree) Right() *Node {
	var parent *Node
	current := tree.Root
	for current != nil {
		parent = current
		current = current.Right
	}
	return parent
}

func (tree *Tree) Floor(key interface{}) (floor *Node, found bool) {
	found = false
	node := tree.Root
	for node != nil {
		compare := tree.Comparator(key, node.Key)
		switch {
		case compare == 0:
			return node, true
		case compare < 0:
			node = node.Left
		case compare > 0:
			floor, found = node, true
			node = node.Right
		}
	}
	if found {
		return floor, true
	}
	return nil, false
}

// Ceiling finds ceiling node of the input key, return the ceiling node or nil if no ceiling is found.
// Second return parameter is true if ceiling was found, otherwise false.
//
// Ceiling node is defined as the smallest node that is larger than or equal to the given node.
// A ceiling node may not be found, either because the tree is empty, or because
// all nodes in the tree are smaller than the given node.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree) Ceiling(key interface{}) (ceiling *Node, found bool) {
	found = false
	node := tree.Root
	for node != nil {
		compare := tree.Comparator(key, node.Key)
		switch {
		case compare == 0:
			return node, true
		case compare < 0:
			ceiling, found = node, true
			node = node.Left
		case compare > 0:
			node = node.Right
		}
	}
	if found {
		return ceiling, true
	}
	return nil, false
}

// Clear removes all nodes from the tree.
func (tree *Tree) Clear() {
	tree.Root = nil
	tree.size = 0
}

// String returns a string representation of container
func (tree *Tree) String() string {
	str := "AVLTree\n"
	output(tree.Root, "", true, &str, false)
	return str
}
func (tree *Tree) StringWithColor() string {
	str := "AVLTree\n"
	output(tree.Root, "", true, &str, true)
	return str
}
func (node *Node) updateHeight() {
	node.height = max(node.Left.Height(), node.Right.Height()) + 1
}
func (node *Node) Height() int {
	if node == nil {
		return 0
	}
	return node.height
}
func (node *Node) String(withColor bool) string {
	if withColor {
		return fmt.Sprintf("%v(%v)", node.Key, node.Height())
	}
	return fmt.Sprintf("%v", node.Key)
}

func output(node *Node, prefix string, isTail bool, str *string, withColor bool) {
	if node == nil {
		return
	}
	if node.Right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.Right, newPrefix, false, str, withColor)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String(withColor) + "\n"
	if node.Left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.Left, newPrefix, true, str, withColor)
	}
}

func (tree *Tree) lookup(key interface{}) *Node {
	node := tree.Root
	for node != nil {
		compare := tree.Comparator(key, node.Key)
		switch {
		case compare == 0:
			return node
		case compare < 0:
			node = node.Left
		case compare > 0:
			node = node.Right
		}
	}
	return nil
}

func (node *Node) maximumNode() *Node {
	if node == nil {
		return nil
	}
	for node.Right != nil {
		node = node.Right
	}
	return node
}
