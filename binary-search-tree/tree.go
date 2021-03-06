package main

import (
	"fmt"
)

// Node defines a single tree node
type Node struct {
	left  *Node
	right *Node
	data  int
}

// BinarySearchTree defines a root node and interface to call function upon
type BinarySearchTree struct {
	root *Node
}

// insert asks a Node to insert a value, recursively
func (n *Node) insert(value int) {
	// if value is smaller than our node's data...
	if value <= n.data {
		// ...then insert the new node to the left
		if n.left == nil { // do we have a left node?
			n.left = &Node{data: value} // if not, make one
		} else {
			// ...otherwise, ask our left node to insert the data recursively
			n.left.insert(value)
		}
	} else {
		// ...otherwise, insert the new node to the right
		if n.right == nil { // do we have a right node?
			n.right = &Node{data: value} // if not, make one
		} else {
			// ...otherwise, ask our right node to insert the data recursively
			n.right.insert(value)
		}
	}
}

// Insert adds a node with a given value to the tree
func (tree *BinarySearchTree) Insert(value int) {
	if tree.root == nil { // do we have a root?
		tree.root = &Node{data: value} // ...if not, create one to start the tree
	} else {
		tree.root.insert(value) // ...if we do have a root, ask it to insert the value
	}
}

// contains walks through the tree starting at a specified Node and determines if it or it's children contain value
func (n *Node) contains(value int) bool {
	// base case; if the current node has the data, then the tree contains the data
	if value == n.data {
		return true
	} else if value < n.data { // if the value is smaller than the current node's data, it should be to the left
		if n.left == nil { // ...but if there is no left node, the tree doesn't contain this value because there's no other nodes to ask
			return false
		} else { // if there is a left node, ask if itself or it's child nodes contain this value
			return n.left.contains(value)
		}
	} else { // ...if value > n.data, then look to the right
		if n.right == nil { // ...but if there is no right node, the tree doesn't contain this value because there's no other nodes to ask
			return false
		} else { // if there is a right node, ask if itself or it's child nodes contain this value
			return n.right.contains(value)
		}
	}
}

// Contains walks through the tree and determines if it includes a certain value
func (tree *BinarySearchTree) Contains(value int) bool {
	return tree.root.contains(value)
}

// Max finds the largest value contained in the tree (or 0 if there is nothing in the tree)
func (tree *BinarySearchTree) Max() int {
	if tree.root == nil { // if the tree is empty, default to a "maximum value" of 0
		return 0
	}

	currentNode := tree.root // start at the root
	for {                    // loop over nodes
		if currentNode.right == nil { // if we are at the right-most node (a leaf)
			return currentNode.data // ...then return the data because it is the largest
		}
		currentNode = currentNode.right // ...otherwise, move further right
	}
}

// Min finds the smallest value contained in the tree (or 0 if there is nothing in the tree)
func (tree *BinarySearchTree) Min() int {
	if tree.root == nil { // if the tree is empty, default to a "minimum value" of 0
		return 0
	}

	currentNode := tree.root // start at the root
	for {                    // loop over nodes
		if currentNode.left == nil { // if we are at the left-most node (a leaf)
			return currentNode.data // ...then return the data because it is the smallest
		}
		currentNode = currentNode.left // ...otherwise, move further left
	}
}

// printTree runs an in-order traversal starting at the specified Node and indentationLevel
func (n *Node) printTree(indentationLevel int) {
	if n != nil { // if there is a node to print
		padding := ""             // stores the indentation characters
		if indentationLevel > 0 { // if we aren't at the root
			if n.left == nil && n.right == nil { // if this is a leaf, don't use a continuation character
				padding += "└"
			} else { // if this is not leaf, use a continuation character
				padding += "├"
			}
		}

		for i := 0; i < indentationLevel; i++ { // assemble left-right padding based on indentationLevel
			padding += "─"
		}

		fmt.Printf("%s%d\n", padding, n.data)
		n.left.printTree(indentationLevel + 1)
		n.right.printTree(indentationLevel + 1)
	}
}

// Print displays a graphical text representation of the tree by traversing it in order
func (tree *BinarySearchTree) Print() {
	tree.root.printTree(0)
}
