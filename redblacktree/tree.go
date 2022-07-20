// References: http://en.wikipedia.org/wiki/Red%E2%80%93black_tree
// https://www.youtube.com/playlist?list=PLpPXw4zFa0uKKhaSz87IowJnOTzh9tiBk
package redblacktree

import (
	"fmt"

	"github.com/emirpasic/gods/utils"
)

type color int64

const (
	black color = iota
	red
)

func (s color) String() string {
	switch s {
	case black:
		return "B"
	}
	return "R"
}

// Tree holds elements of the red-black tree
type Tree struct {
	Root       *Node
	size       int
	Comparator utils.Comparator
}

// Node is a single element within the tree
type Node struct {
	Key    interface{}
	Value  interface{}
	color  color
	Left   *Node
	Right  *Node
	Parent *Node
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

// Put inserts node into the tree.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree) Put(key interface{}, value interface{}) {
	n := &Node{Key: key, Value: value, color: red}
	var insertNode *Node
	if tree.Root == nil {
		tree.Root = n
		insertNode = n
		n.color = black
	} else {
		loop := true
		parent := tree.Root
		for loop {
			compare := tree.Comparator(n.Value, parent.Value)
			switch {
			case compare == 0:
				parent.Value = value
				return
			case compare < 0:
				if parent.Left == nil {
					loop = false
					parent.Left = n
					n.Parent = parent
					insertNode = n
				} else {
					parent = parent.Left
				}
			case compare > 0:
				if parent.Right == nil {
					loop = false
					parent.Right = n
					n.Parent = parent
					insertNode = n

				} else {
					parent = parent.Right
				}

			}

		}
	}
	tree.insertCase1(insertNode)
	tree.size = tree.size + 1
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

// Remove remove the node from the tree by key.
// Key should adhere to the comparator's type assertion, otherwise method panics.
func (tree *Tree) Remove(key interface{}) {
	var child *Node
	node := tree.lookup(key)
	if node == nil {
		return
	}
	if node.Left != nil && node.Right != nil {
		pred := node.Left.maximumNode()
		node.Key = pred.Key
		node.Value = pred.Value
		node = pred
	}
	if node.Left == nil || node.Right == nil {
		if node.Right == nil {
			child = node.Left
		} else {
			child = node.Right
		}
		if node.color == black {
			node.color = nodeColor(child)
			tree.deleteCase1(node)
		}
		tree.replaceNode(node, child)
		if node.Parent == nil && child != nil {
			child.color = black
		}
	}
	tree.size--
}

// Empty returns true if tree does not contain any nodes
func (tree *Tree) Empty() bool {
	return tree.size == 0
}

// Size returns number of nodes in the tree.
func (tree *Tree) Size() int {
	return tree.size
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

// // Keys returns all keys in-order
// func (tree *Tree) Keys() []interface{} {
// 	keys := make([]interface{}, tree.size)
// 	it := tree.Iterator()
// 	for i := 0; it.Next(); i++ {
// 		keys[i] = it.Key()
// 	}
// 	return keys
// }

// // Values returns all values in-order based on the key.
// func (tree *Tree) Values() []interface{} {
// 	values := make([]interface{}, tree.size)
// 	it := tree.Iterator()
// 	for i := 0; it.Next(); i++ {
// 		values[i] = it.Value()
// 	}
// 	return values
// }

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

// Floor Finds floor node of the input key, return the floor node or nil if no floor is found.
// Second return parameter is true if floor was found, otherwise false.
//
// Floor node is defined as the largest node that is smaller than or equal to the given node.
// A floor node may not be found, either because the tree is empty, or because
// all nodes in the tree are larger than the given node.
//
// Key should adhere to the comparator's type assertion, otherwise method panics.
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
	str := "RedBlackTree\n"
	if !tree.Empty() {
		output(tree.Root, "", true, &str, false)
	}
	return str
}
func (tree *Tree) StringWithColor() string {
	str := "RedBlackTree\n"
	if !tree.Empty() {
		output(tree.Root, "", true, &str, true)
	}
	return str
}

func (node *Node) String(withColor bool) string {
	if withColor {
		return fmt.Sprintf("%v(%v)", node.Key, node.color)
	}
	return fmt.Sprintf("%v", node.Key)
}

func output(node *Node, prefix string, isTail bool, str *string, withColor bool) {
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

func (node *Node) grandparent() *Node {
	if node != nil && node.Parent != nil {
		return node.Parent.Parent
	}
	return nil
}

func (node *Node) uncle() *Node {
	if node == nil || node.Parent == nil || node.Parent.Parent == nil {
		return nil
	}
	return node.Parent.sibling()
}

func (node *Node) sibling() *Node {
	if node == nil || node.Parent == nil {
		return nil
	}
	if node == node.Parent.Left {
		return node.Parent.Right
	}
	return node.Parent.Left
}

func (tree *Tree) rotateLeft(node *Node) {
	child := node.Right
	tree.replaceNode(node, child)
	node.Right = child.Left
	if child.Left != nil {
		child.Left.Parent = node
	}
	child.Left = node
	node.Parent = child
}

func (tree *Tree) rotateRight(node *Node) {
	child := node.Left
	tree.replaceNode(node, child)
	node.Left = child.Right
	if child.Right != nil {
		child.Right.Parent = node
	}
	child.Right = node
	node.Parent = child
}

func (tree *Tree) replaceNode(old *Node, new *Node) {
	parent := old.Parent
	if parent != nil {
		if parent.Left == old {
			parent.Left = new
		} else {
			parent.Right = new
		}
	} else {
		tree.Root = new
	}
	new.Parent = parent
}

func (tree *Tree) insertCase1(node *Node) {
	if node.Parent == nil {
		node.color = black
		return
	}
	if node.Parent.color == black {
		return
	}
	tree.insertRedNode(node)
}
func (tree *Tree) insertRedNode(node *Node) {
	uncle := node.uncle()
	if uncle == nil || uncle.color == black {
		tree.rotate(node)
	} else {
		//flip Color
		uncle.color = black
		node.Parent.color = black
		grandparent := node.Parent
		grandparent.color = red
		tree.insertCase1(grandparent)
	}
}
func (tree *Tree) rotate(node *Node) {
	grandparent := node.grandparent()
	parent := node.Parent
	if parent.Left == node && grandparent.Right == parent {
		tree.rotateRight(parent)
		node = parent
	} else if parent.Right == node && grandparent.Left == parent {
		tree.rotateLeft(parent)
		node = parent
	}
	tree.rotate2(node)
}
func (tree *Tree) rotate2(node *Node) {
	grandparent := node.grandparent()
	parent := node.Parent
	if parent.Left == node && grandparent.Left == parent {
		tree.rotateRight(grandparent)
	} else if parent.Right == node && grandparent.Right == parent {
		tree.rotateLeft(grandparent)
	}
	parent.color = black
	grandparent.color = red
	// node.color = red
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

func (tree *Tree) deleteCase1(node *Node) {
	if node.Parent == nil {
		return
	}
	tree.deleteCase2(node)
}

func (tree *Tree) deleteCase2(node *Node) {
	sibling := node.sibling()
	if nodeColor(sibling) == red {
		node.Parent.color = red
		sibling.color = black
		if node == node.Parent.Left {
			tree.rotateLeft(node.Parent)
		} else {
			tree.rotateRight(node.Parent)
		}
	}
	tree.deleteCase3(node)
}

func (tree *Tree) deleteCase3(node *Node) {
	sibling := node.sibling()
	if nodeColor(node.Parent) == black &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Left) == black &&
		nodeColor(sibling.Right) == black {
		sibling.color = red
		tree.deleteCase1(node.Parent)
	} else {
		tree.deleteCase4(node)
	}
}

func (tree *Tree) deleteCase4(node *Node) {
	sibling := node.sibling()
	if nodeColor(node.Parent) == red &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Left) == black &&
		nodeColor(sibling.Right) == black {
		sibling.color = red
		node.Parent.color = black
	} else {
		tree.deleteCase5(node)
	}
}

func (tree *Tree) deleteCase5(node *Node) {
	sibling := node.sibling()
	if node == node.Parent.Left &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Left) == red &&
		nodeColor(sibling.Right) == black {
		sibling.color = red
		sibling.Left.color = black
		tree.rotateRight(sibling)
	} else if node == node.Parent.Right &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Right) == red &&
		nodeColor(sibling.Left) == black {
		sibling.color = red
		sibling.Right.color = black
		tree.rotateLeft(sibling)
	}
	tree.deleteCase6(node)
}

func (tree *Tree) deleteCase6(node *Node) {
	sibling := node.sibling()
	sibling.color = nodeColor(node.Parent)
	node.Parent.color = black
	if node == node.Parent.Left && nodeColor(sibling.Right) == red {
		sibling.Right.color = black
		tree.rotateLeft(node.Parent)
	} else if nodeColor(sibling.Left) == red {
		sibling.Left.color = black
		tree.rotateRight(node.Parent)
	}
}

func nodeColor(node *Node) color {
	if node == nil {
		return black
	}
	return node.color
}
