package main

import (
	"workbench/workbench"
)

func main() {
	ll := &workbench.LinkedList{}
	ll.Insert(10)
	ll.Insert(20)
	ll.Insert(30)
	ll.Insert(40)

	ll.Print()
}
