package main

import (
	"fmt"
	"math/rand"
)

func QuickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}

	pivot := list[rand.Intn(len(list))]
	var greaterThan, sameAs, lessThan []int

	for _, v := range list {
		if v > pivot {
			greaterThan = append(greaterThan, v)
		} else if v < pivot {
			lessThan = append(lessThan, v)
		} else {
			sameAs = append(sameAs, v)
		}
	}

	less := QuickSort(lessThan)
	greater := QuickSort(greaterThan)
	sorted := append(less, sameAs...)
	sorted = append(sorted, greater...)

	return sorted
}

func main() {
	got := QuickSort([]int{5, 1, 1, 2, 0, 0})
	fmt.Printf("got: %d\n", got)

}
