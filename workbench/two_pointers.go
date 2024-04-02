package workbench

func TwoSumUsingTwoPointers(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1
	for left < right {
		current := nums[left] + nums[right]
		if current == target {
			return true
		}
		if current > target {
			right--
		} else {
			left++
		}
	}

	return false
}
