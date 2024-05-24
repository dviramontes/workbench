package workbench

func BinarySearchRecursive(list []int, target int) int {
	low := 0
	high := len(list) - 1

	return binarySearchRecursive(list, target, low, high)
}

func binarySearchRecursive(list []int, target, low, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2
	guess := list[mid]

	if guess == target { // base case
		return mid
	}

	if guess > target { // recursive case
		return binarySearchRecursive(list, target, low, mid-1)
	} else {
		return binarySearchRecursive(list, target, mid+1, high)
	}
}
