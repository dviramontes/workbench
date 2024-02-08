package workbench

func RemoveElementsFromSortedArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	uniqueIndex := 1
	for i := 1; i < len(nums); i++ {
		// if the current element is not a duplicate
		// move it to the next position of the unique index
		if nums[i] != nums[i-1] {
			nums[uniqueIndex] = nums[i]
			uniqueIndex++
		}
	}

	return uniqueIndex
}
