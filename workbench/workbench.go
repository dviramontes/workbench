package workbench

import (
	"bytes"
	"sort"
	"strings"
)

func MaxMeetings(start []int, end []int) []int {
	// keep track of the meetings that can be scheduled
	var meetings []int

	// sort start of meetings by end time
	sort.Slice(start, func(i, j int) bool {
		return end[i] < end[j]
	})

	finishCursor := end[0]
	// we know that at least the first meeting can be scheduled
	meetings = append(meetings, 1)

	for i := 0; i < len(start); i++ {
		// if the current meeting starts after the last finishCursor, then schedule it
		if start[i] >= finishCursor {
			meetings = append(meetings, i+1)
			finishCursor = end[i]
		}
	}

	return meetings
}
func CanAttendMeetings(intervals []Interval) bool {
	var canAttend bool = true
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	firstMeeting := intervals[0]

	for i := 1; i < len(intervals); i++ {
		if firstMeeting.End > intervals[i].Start {
			canAttend = false
		}
	}

	return canAttend
}

func TwoSum(nums []int, target int) []int {
	compliments := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		compliment := target - nums[i]
		if j, found := compliments[compliment]; found {
			return []int{j, i}
		}
		compliments[nums[i]] = i
	}

	return []int{}
}

func MergeAlternate(word1 string, word2 string) string {
	var buf bytes.Buffer

	if len(word1) == len(word2) {
		for i := 0; i < len(word1); i++ {
			buf.WriteRune(rune(word1[i]))
			buf.WriteRune(rune(word2[i]))
		}
		return buf.String()
	}

	var shortWordLength int
	var rest string
	if len(word1) > len(word2) {
		shortWordLength = len(word2)
		rest = word1[len(word2):]
	} else {
		shortWordLength = len(word1)
		rest = word2[len(word1):]
	}

	for i := 0; i < shortWordLength; i++ {
		buf.WriteRune(rune(word1[i]))
		buf.WriteRune(rune(word2[i]))
	}

	return buf.String() + rest
}

// GreatestCommonDivisor
// time complexity: min(m, n) * (n * m)
func GreatestCommonDivisor(str1 string, str2 string) string {
	if !strings.Contains(str1, str2) {
		return ""
	}

	l1 := len(str1)
	l2 := len(str2)
	minLength := min(l1, l2)
	for l := minLength; l >= minLength; l-- {
		if isDivisor(l, l1, l2) {
			return str1[:l]
		}
	}

	return ""
}

func isDivisor(l, l1, l2 int) bool {
	if l1%l == 0 || l2%l == 0 {
		return false
	}
	return true
}

// ContainsDuplicate
// brute force
func ContainsDuplicate(n []int) bool {
	sort.Slice(n, func(i, j int) bool {
		return n[i] < n[j]
	})

	for i := 0; i < len(n)-1; i++ {
		if n[i] == n[i+1] {
			return true
		}
	}
	return false
}

// ContainsDuplicateBetter
func ContainsDuplicateBetter(n []int) bool {
	m := make(map[int]bool)
	for _, v := range n {
		if _, found := m[v]; found {
			return true
		}
		m[v] = true
	}
	return false
}
