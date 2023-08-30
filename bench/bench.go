package bench

import (
	"bytes"
	"sort"
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
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
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
		shortWordLength = len(word1) - len(word2)
		rest = word1[len(word2):]
	} else {
		shortWordLength = len(word2) - len(word1)
		rest = word2[len(word1):]
	}

	for i := 0; i < shortWordLength; i++ {
		buf.WriteRune(rune(word1[i]))
		buf.WriteRune(rune(word2[i]))
	}

	return buf.String() + rest
}
