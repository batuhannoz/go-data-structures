package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

func (s *Stack) Pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func main() {
	myStack := &Stack{items: []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(101)
	fmt.Println(myStack)
	pop := myStack.Pop()
	fmt.Printf("%d\t%v", pop, myStack)
}
