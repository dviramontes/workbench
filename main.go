package main

import (
	"fmt"
	"workbench/workbench"
)

func main() {
	qs := workbench.QuickSort([]int{0, 1, 3, 2})
	fmt.Println(qs)
}
