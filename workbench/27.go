package workbench

// RemoveElementFromArray
// https://leetcode.com/problems/remove-element/
func RemoveElementFromArray(nums []int, val int) int {
	insertAtIndex := 0

	for _, n := range nums {
		if n != val {
			nums[insertAtIndex] = n
			insertAtIndex++
		}
	}

	return insertAtIndex
}
