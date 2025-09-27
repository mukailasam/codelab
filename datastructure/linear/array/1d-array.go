package main

import "fmt"

func main() {
	// define an array of type string with length of 4
	arr := [4]string{"Go", "C++", "Rust", "Python"}
	// Problem: access the element "Go" in the above array
	// solution: arr[0]
	// Note: the index indicate the position of an element in an array.
	// So, arr[0] indicate the first position of the array
	// Here we are using single index because it a on-dimensional
	// array
	// see arr[0] as go to the first position and get the element
	// in that position, which is "Go"
	fmt.Println(arr[0])
}
