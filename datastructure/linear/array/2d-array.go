package main

import "fmt"

func main() {
	// Note: 2D array are arranged in form of rows and columns
	// multArr[3][3]int specify the kind of element and the
	// number of rows and column of this type of element it can
	// hold
	multiArr := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// Problem: access the element "8", in the above array
	// solution: multiArr[2][1]
	// Note: the first index indicate the row position
	// of the element of the array we want to access, while
	// the second index indicate the column position
	// of the element of the array we want to access.
	// So, the two indexes indicate the position of an
	// element of a two-dimensional array
	// Here the first index is 2 which indicate the third row and
	// the second index is 1 which indicate the second column of
	// the third row. so, see multiArr[2][1] as go to the the
	// third row, in the third row get me the element in the
	// second column, which is 8
	fmt.Println(multiArr[2][1])
}
