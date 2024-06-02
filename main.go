package main

import "fmt"

func BinarySearch(list []int, target int) int {
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (high + low) / 2
		guess := list[mid]
		if target == guess {
			return mid
		}

		if guess > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func main() {
	got := BinarySearch([]int{1, 2, 3, 4, 5}, 3)
	fmt.Printf("got: %d\n", got)

}
