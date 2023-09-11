package main

import (
	"fmt"
	"workbench/workbench"
)

func main() {
	// find the maximum number of meetings
	//s := []int{1, 3, 0, 5, 8, 5}
	//f := []int{2, 4, 6, 7, 9, 9}

	//result := bench.MaxMeetings(s, f)
	//fmt.Printf("meetings: %+v\n", result)

	// can attend meetings, that don't overlap
	//i1 := []bench.Interval{{5, 10}, {0, 30}, {15, 20}}
	//i2 := []Interval{{5, 8}, {9, 15}}
	//canAttend := bench.CanAttendMeetings(i1)
	//fmt.Printf("can attent meetings: %+v\n", canAttend)

	// two sum
	// Input: nums = [2,7,11,15], target = 9
	// Output: [0,1]
	// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
	//nums := []int{2, 7, 11, 15}
	//target := 9
	//result = bench.TwoSum(nums, target)
	//fmt.Printf("two sum: %+v\n", result)

	// Merge Alternate
	// Input: word1 = "abc", word2 = "pqr"
	// Output: "apbqcr"
	// Explanation: The merged string will be merged as so:
	// word1:  a   b   c
	// word2:    p   q   r
	// merged: a p b q c r
	//word1 := "abc"
	//word2 := "pqr"
	//mergeAlt := workbench.MergeAlternate(word1, word2)
	//fmt.Printf("merge alternate: %+v\n", mergeAlt)

	// 1071. Greatest Common Divisor of Strings
	// Input: str1 = "ABCABC", str2 = "ABC"
	// Output: "ABC"
	// Input: str1 = "LEET", str2 = "CODE"
	// Output: ""
	gcd := workbench.GreatestCommonDivisor("ABCABC", "ABC")
	fmt.Printf("gcd: %+v\n", gcd)
}
