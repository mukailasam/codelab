// a simple binary tree
// Note:
// although the code explain itself
// The comments are for beginners that are still
// struggling with golang syntax

package main

import "fmt"

// define a tree
type Tree struct {
	rootNode *Node
}

// define a node
type Node struct {
	value     int
	leftNode  *Node
	rightNode *Node
}

func main() {
	// create a tree
	tree := Tree{}

	// create the root node and set it value to 4
	tree.rootNode = &Node{value: 4}

	// create the right child node for the root node
	// and set it value to 5
	tree.rootNode.rightNode = &Node{value: 5}

	// create the left child node for the root node
	// and set it value to 3

	tree.rootNode.leftNode = &Node{value: 3}
	// create the right child node for the left child
	// node of the root node and set it value to 20

	tree.rootNode.leftNode.rightNode = &Node{value: 20}
	// create a left child node for the left child node
	// of the root node and set the value to 1

	tree.rootNode.leftNode.leftNode = &Node{value: 1}

	// called to perform in-order traversal
	tree.rootNode.inOrderTraversal()
}

// traverse the tree using in-order traversal and print
// on the console each value of every node in the tree
func (t *Node) inOrderTraversal() {
	if t == nil {
		return
	}

	t.leftNode.inOrderTraversal()
	fmt.Printf("%d ", t.value)
	t.rightNode.inOrderTraversal()
}
