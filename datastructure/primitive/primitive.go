package main

import (
	"fmt"
)

func main() {
	// define some variable.
	// In Go, there is no char data type instead
	// Go uses byte and rune to represent character values.
	var integer int = 10
	var character rune = 'S'
	var decimal float64 = 3.141592653589793238462643383279502884197
	var boolean bool = true
	// Print out to the console each value stored in each
	// one of the variable
	fmt.Println("Integer:", integer)
	fmt.Println("Character:", character)
	fmt.Println("Float:", decimal)
	fmt.Println("Boolean:", boolean)
}
