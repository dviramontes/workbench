package workbench

// MergeSortedArray
// https://leetcode.com/problems/merge-sorted-array/
func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) []int {
	// merge from the end of the array
	i := m - 1
	j := n - 1
	cursor := m + n - 1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[cursor] = nums1[i]
			i--
		} else {
			nums1[cursor] = nums2[j]
			j--
		}
		cursor--
	}

	// if there are remaining elements in nums2
	for j >= 0 {
		nums1[cursor] = nums2[j]
		cursor--
		j--
	}

	return nums1
}
