package main

import "fmt"

const arraySize = 7

type hashTable struct {
	array [arraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

type bucketNode struct {
	key   string
	value any
	next  *bucketNode
}

func main() {
	hashTable := Init() // Initialize the hash table
	// Insert a key-value pair into the table
	hashTable.insert("newt", []string{"1", "2", "3"})
	// Retrieve and print the value associated with the key "newt".
	fmt.Println(hashTable.get("newt"))
	// Delete the key "newt" from the hash table
	hashTable.delete("newt")
	// Try to retrieve the value again after deletion
	fmt.Println(hashTable.get("newt"))
}

// Init initializes the hash table with empty buckets
func Init() *hashTable {
	result := hashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return &result
}

// insert inserts a key-value pair into the hash table
func (h *hashTable) insert(key string, value any) {
	index := hash(key)
	h.array[index].insert(key, value)
}

// delete removes a key from the hash table
func (h *hashTable) delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// get retrieves a value from the hash table based on the key
func (h *hashTable) get(key string) any {
	index := hash(key)
	return h.array[index].get(key)
}

func (h *hashTable) search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// insert inserts a new node at the beginning of the bucket's linked list
func (b *bucket) insert(key string, value any) {
	if !b.search(key) {
		newNode := &bucketNode{key: key, value: value}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(key, "already exists")
	}
}

// delete removes a node from the linked list if the key matches
func (b *bucket) delete(key string) {
	if b.head.key == key {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == key {
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

// get searches the bucket and returns the value or a not-found message
func (b *bucket) get(key string) any {
	if b.search(key) {
		currentNode := b.head
		for currentNode != nil {
			if currentNode.key == key {
				return currentNode.value
			}
			currentNode = currentNode.next
		}
		return fmt.Sprintf("%v not exists", key)
	} else {
		return fmt.Sprintf("%v not exists", key)
	}
}

// search checks if a node with the key exists in the bucket
func (b *bucket) search(key string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// hash is a simple hash function that returns the sums of the ASCII values of the key's characters divided by arraySize
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % arraySize // Modulo to stay within the table bounds
}
