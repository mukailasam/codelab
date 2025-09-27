package main

// A rooted Tree
// Binary search Tree
import "fmt"

// define node
type Node struct {
	value     int
	leftNode  *Node
	rightNode *Node
}

func main() {
	// start Binary Search Tree
	// set the root node
	bst := NewBSTree(4)
	// Insert node into the tree
	bst.insertANode(5)
	bst.insertANode(3)
	bst.insertANode(20)
	bst.insertANode(1)
	bst.preOrderTraversal()
	fmt.Println()
	bst.inOrderTraversal()
	fmt.Println()
	bst.postOrderTraversal()
	fmt.Println()
}
func NewBSTree(rootNodeValue int) *Node {
	bst := &Node{value: rootNodeValue}
	return bst
}
func (bt *Node) insertANode(value int) {
	if value <= bt.value {
		if bt.leftNode != nil {
			bt.leftNode.insertANode(value)
			return
		} else {
			bt.leftNode = &Node{value: value}
			return
		}
	} else {
		if bt.rightNode != nil {
			bt.rightNode.insertANode(value)
			return
		} else {
			bt.rightNode = &Node{value: value}
			return
		}
	}
}
func (bt *Node) inOrderTraversal() {
	if bt == nil {
		return
	}
	bt.leftNode.inOrderTraversal()
	fmt.Printf("%d ", bt.value)
	bt.rightNode.inOrderTraversal()
}
func (bt *Node) preOrderTraversal() {
	if bt == nil {
		return
	}
	fmt.Printf("%d ", bt.value)
	bt.leftNode.preOrderTraversal()
	bt.rightNode.preOrderTraversal()
}
func (bt *Node) postOrderTraversal() {
	if bt == nil {
		return
	}
	bt.leftNode.preOrderTraversal()
	bt.rightNode.preOrderTraversal()
	fmt.Printf("%d ", bt.value)
}
