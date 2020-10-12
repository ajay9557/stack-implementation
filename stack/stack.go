package stack

import (
	"fmt"
)

// Stack will store the string value
var Stack []string

//Max used to set the limit of stack
var Max = 10

//Push will insert the string value into the stack
func Push(value string) []string {
	if len(Stack) == Max {
		fmt.Println("stack is full")
		return Stack
	}
	Stack = append(Stack, value)
	fmt.Println("successfully pushed", Stack)
	return Stack
}

//Pop will delete the top	 value from the stack
func Pop() []string {
	if len(Stack) <= 0 {
		fmt.Println("Stack is empty")
		return Stack
	}
	n := len(Stack) - 1
	fmt.Println("successfully poped", Stack[n])
	Stack = Stack[:n]
	return Stack
}
