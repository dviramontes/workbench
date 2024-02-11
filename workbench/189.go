package workbench

import (
	"golang.org/x/exp/slices"
)

func RotateArrayNTimes(nums []int, k int) []int {
	l := len(nums)
	k = k % l

	slices.Reverse(nums[:l]) // reverse the whole array
	slices.Reverse(nums[:k]) // reverse the first k elements
	slices.Reverse(nums[k:]) // reverse the rest of the elements
	return nums
}
