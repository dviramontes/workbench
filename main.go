package main

import (
	"fmt"
	"sort"
)

func maxMeetings(start []int, end []int) []int {
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

func main() {
	// find the maximum number of meetings
	s := []int{1, 3, 0, 5, 8, 5}
	f := []int{2, 4, 6, 7, 9, 9}

	result := maxMeetings(s, f)
	fmt.Printf("meetings: %+v\n", result)

	// can attend meetings, that don't overlap
	i1 := []Interval{{5, 10}, {0, 30}, {15, 20}}
	//i2 := []Interval{{5, 8}, {9, 15}}
	canAttend := canAttendMeetings(i1)
	fmt.Printf("can attent meetings: %+v\n", canAttend)
}

func canAttendMeetings(intervals []Interval) bool {
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

type Interval struct {
	Start, End int
}
