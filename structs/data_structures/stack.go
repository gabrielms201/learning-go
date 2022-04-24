package stack

import "fmt"

const (
	size = 100
)

func GetSize() int { return size }

type TYPE int
type Stack struct {
	data [size]TYPE
	top  int
}

func NewStack() *Stack {
	st := Stack{}
	st.top = -1
	return &st
}

// We need to start the atributes with uppercase to made it publicly accessible

func (st *Stack) Pop() TYPE {
	if st.IsEmpty() {
		fmt.Println("Stack empty")
		return 0;
	}
	var elem TYPE = st.data[st.top]
	st.top --
	return elem

}

func (st *Stack) Push(element TYPE) bool {
	if st.IsFull() {
		fmt.Println("ERROR: STACK OVERFLOW")
		return false
	}
	st.top ++
	st.data[st.top] = element
	return true
}
func (st *Stack) Peek() TYPE {
	if st.IsEmpty(){
		fmt.Println("Stack Is Empty")
		return 0
	}
	return st.data[st.top]
}
func (st *Stack) IsEmpty() bool { return st.top == -1 }
func (st *Stack) IsFull() bool  { return st.top == size-1 }
