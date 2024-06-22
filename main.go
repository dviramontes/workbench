package main

import (
	"fmt"
	"math/rand"
)

func QuickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}

	r := rand.Intn(len(list))
	pivot := list[r]
	var sameAs, lessThan, greaterThan []int
	for _, v := range list {
		if v > pivot {
			greaterThan = append(greaterThan, v)
		} else if v < pivot {
			lessThan = append(lessThan, v)
		} else {
			sameAs = append(sameAs, v)
		}
	}

	sorted := append(QuickSort(lessThan), sameAs...)
	sorted = append(sorted, QuickSort(greaterThan)...)

	return sorted
}

func main() {
	got := QuickSort([]int{12, 2, 6, 23, 55, 100, 5, 2, 33, 444})
	fmt.Printf("got: %d\n", got)
}
