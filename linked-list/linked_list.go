package main

import "fmt"

type node struct {
	data string
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d\t", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteWithValue(value string) {
	if l.length == 0 {
		fmt.Println("The list is empty")
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func main() {
	mylist := linkedList{}
	mylist.prepend(&node{data: "Batuhan"})
	mylist.prepend(&node{data: "Ozdemir"})
	mylist.prepend(&node{data: "Luke"})
	mylist.prepend(&node{data: "Skywalker"})
	mylist.prepend(&node{data: "Harry"})
	mylist.prepend(&node{data: "Potter"})

	mylist.printListData()
	mylist.deleteWithValue("Batuhan")
	mylist.printListData()
	emptyList := linkedList{}
	emptyList.printListData()
}
