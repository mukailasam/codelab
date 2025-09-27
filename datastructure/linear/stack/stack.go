package main

import (
	"fmt"
	"strconv"
)

type Element struct {
	elementValue int
}

type Stack struct {
	elements     []*Element
	elementCount int
}

func main() {
	elem1 := &Element{5}
	elem2 := &Element{4}
	elem3 := &Element{7}
	elem4 := &Element{6}
	s := &Stack{}
	s.New()
	s.Push(elem1)
	s.Push(elem2)
	s.Push(elem3)
	s.Push(elem4)
	s.Pop()
	s.Pop()
	fmt.Println(s.elements)
}

func (element Element) String() string {
	return strconv.Itoa(element.elementValue)
}

func (stack *Stack) New() {
	stack.elements = make([]*Element, 0)
}

func (start *Stack) Push(element *Element) {
	start.elements = append(start.elements, element)
	start.elementCount += 1
}

func (stack *Stack) Pop() []*Element {
	if stack.elementCount == 0 {
		return nil
	}
	length := len(stack.elements)
	stack.elements = stack.elements[:length-1]
	stack.elementCount -= 1
	return stack.elements
}
