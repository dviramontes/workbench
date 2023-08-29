package main

import (
	"fmt"
	"workbench/bench"
)

func main() {
	// find the maximum number of meetings
	s := []int{1, 3, 0, 5, 8, 5}
	f := []int{2, 4, 6, 7, 9, 9}

	result := bench.MaxMeetings(s, f)
	fmt.Printf("meetings: %+v\n", result)

	// can attend meetings, that don't overlap
	i1 := []bench.Interval{{5, 10}, {0, 30}, {15, 20}}
	//i2 := []Interval{{5, 8}, {9, 15}}
	canAttend := bench.CanAttendMeetings(i1)
	fmt.Printf("can attent meetings: %+v\n", canAttend)

	// two sum
	// Input: nums = [2,7,11,15], target = 9
	// Output: [0,1]
	// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
	nums := []int{2, 7, 11, 15}
	target := 9
	result = bench.TwoSum(nums, target)
	fmt.Printf("two sum: %+v\n", result)
}
