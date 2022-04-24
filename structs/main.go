package main

import (
	"fmt"
	. "structures/data_structures"
)

func main() {
	stack := NewStack()
	stack.Push(1)   // 1
	stack.Push(2)   // 1, 2
	stack.Pop()     // 1
	stack.Push(3)   // 1, 3
	stack.Push(4)   // 1, 3, 4
	stack.Push(5)   // 1, 3, 4, 5
	stack.Push(6)   // 1, 3, 4, 5, 6
	stack.Pop()     // 1, 3, 4, 5,
	stack.Push(7)   // 1, 3, 4, 5, 7
	stack.Push(8)   // 1, 3, 4, 5, 7, 8
	for !stack.IsEmpty(){
		fmt.Println(stack.Pop())
	}
	// 8, 7, 5, 4, 3 1

}
