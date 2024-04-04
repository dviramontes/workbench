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

func CombineArrayUsingTwoPointers(a []int, b []int) (c []int) {
	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}

	// ensure both iterables are exhausted
	for i < len(a) {
		c = append(c, a[i])
		i++
	}

	for j < len(b) {
		c = append(c, b[j])
		j++
	}
	return c
}

func IsSubSequenceUsingTwoPointers(a string, b string) bool {
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
		} else {
			j++
		}
	}

	return i == len(a)
}
