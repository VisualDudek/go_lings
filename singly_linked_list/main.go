package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Insert(value int) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (l *LinkedList) Display() {
	current := l.head

	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}

	fmt.Println("nil")
}

func (l *LinkedList) ToSlice() []int {
	current := l.head
	res := []int{}

	for current != nil {
		res = append(res, current.data)
		current = current.next
	}

	return res
}

func (l *LinkedList) InsertFromSlice(s []int) {
	for _, v := range s {
		l.Insert(v)
	}
}

func main() {
	list := &LinkedList{}

	list.Insert(1)
	list.Insert(2)
	list.Insert(3)

	list.Display()

	fmt.Println(list.ToSlice())
}
