package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myStack := &Queue{items: []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}}
	fmt.Println(myStack)
	myStack.Enqueue(100)
	myStack.Enqueue(101)
	fmt.Println(myStack)
	pop := myStack.Dequeue()
	fmt.Printf("%d\t%v", pop, myStack)
}
