package workbench

import (
	"fmt"
)

// LinkedList is a singly linked list
// with a head and tail pointer
// and a length counter
// it has methods to add and remove nodes
// and to print the list

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	head   *Node
	length int // because of go's default values, we don't have to initialize to 0
}

func (ll *LinkedList) Insert(data int) {
	newNode := &Node{Data: data}

	if ll.head == nil {
		ll.head = newNode
		ll.length++
		return
	}

	current := ll.head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	ll.length++
}

func (ll *LinkedList) Print() {
	current := ll.head
	fmt.Printf("length -> %d\n", ll.length)

	for current != nil {
		fmt.Printf("-> %d\n", current.Data)
		current = current.Next
	}
}

func (ll *LinkedList) Search(data int) bool {
	current := ll.head
	for current != nil {
		if current.Next.Data == data {
			return true
		}
		current = current.Next
	}

	return false
}
