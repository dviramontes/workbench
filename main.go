package main

import (
	"fmt"
	"workbench/workbench"
)

func main() {
	set := workbench.Set{}
	set.Add(1)
	set.Add(2)
	set.Remove(3)

	fmt.Printf("set size: %d\n", set.Size())
}
