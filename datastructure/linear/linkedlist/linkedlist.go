package main

// Linkedlist DS in golang

import (
	"fmt"
)

// define the node
type Node struct {
	Value int   // element of the node
	Next  *Node // pointer to the next node
}

var root *Node // the root of the node
var size int

// entry point
func main() {
	addNodeAtEnd(root, 1)
	addNodeAtEnd(root, 2)
	addNodeAtEnd(root, 3)
	addNodeAtEnd(root, 4)
	addNodeAtFront(root, 10)
	addNodeAtAny(root, 3, 44)
	delNodeAtFront(root)
	delNodeAtEnd(root)
	delNodeAtAny(root, 2)
	fmt.Println(lookupNode(root, 44))
	fmt.Println(Lsize(root))
	traverse(root)
	reverse(root)
}

// it adds node to the end of the linked list
func addNodeAtEnd(t *Node, v int) int {
	if root == nil {
		t = &Node{v, nil}
		root = t
		size += 1
		return 0
	}
	if v == t.Value {
		fmt.Println("Node already exits")
		return -1
	}
	if t.Next == nil {
		t.Next = &Node{v, nil}
		size += 1
		return -2
	}
	return addNodeAtEnd(t.Next, v)
}

// it adds node to the front of the linked list
func addNodeAtFront(t *Node, val int) {
	if t == nil {
		fmt.Println("Empty List ")
	}
	t = &Node{val, nil}
	t.Next = root
	root = t
}

// return the size of the linked list
func Lsize(t *Node) int {
	if t == nil {
		fmt.Println("-> Empty List")
		return 0
	}
	i := 0
	for t != nil {
		i += 1
		t = t.Next
	}
	return i
}

// traverse the node and output each element of the node
func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty List")
		return
	}
	for t != nil {
		fmt.Print(t.Value, " -> ")
		t = t.Next
	}
	print("\n")
}

// look for a specific elements in the nodes
// return true if element found and false if not
func lookupNode(t *Node, v int) bool {
	if root == nil {
		t := &Node{v, nil}
		root = t
		return false
	}
	if v == t.Value {
		return true
	}
	if t.Next == nil {
		return false
	}
	return lookupNode(t.Next, v)
}

// adds node at any point/place/position/location in the linked list
func addNodeAtAny(t *Node, pos, v int) {
	if t == nil {
		fmt.Println("Emty List")
	}
	abs := 1
	index := 2
	for t.Next != nil {
		temp := t.Next
		oldt := t
		if pos == abs {
			t = &Node{v, oldt}
			root = t
			return
		}
		if pos == index {
			t.Next = &Node{v, temp}
			return
		}
		t = t.Next
		index += 1
	}
}

// reverse the linked list
// it goes through each of the node and output the element
// in reverse order starting from the end of the linked list
// to thefront of the linked list
func reverse(t *Node) {
	if t == nil {
		return
	}
	reverse(t.Next)
	fmt.Printf("%d -> ", t.Value)
}

// deletes the node at the end of the linked list
func delNodeAtEnd(t *Node) {
	if t == nil {
		fmt.Println("Empty List")
	}
	temp := t
	for t.Next != nil {
		temp = t
		t = t.Next
	}
	temp = temp
	temp.Next = nil
}

// deletes the node at the front of the linked list
func delNodeAtFront(t *Node) {
	if t == nil {
		fmt.Println("Empty List")
	}
	temp := t.Next
	t = nil
	root = temp
}

// delete a node at any point/place/location/position in the linked list
func delNodeAtAny(t *Node, pos int) {
	if t == nil {
		fmt.Println(" Empty List")
	}
	i := 0
	temp := t
	for t != nil {
		i += 1
		if pos == 1 {
			temp = t.Next
			root = temp
			tout := t
			tout = tout
			tout = nil
			return
		}
		if pos == i {
			temp.Next = t.Next
			t = nil
			return
		}
		//temp = temp
		temp = t
		t = t.Next
	}
}
