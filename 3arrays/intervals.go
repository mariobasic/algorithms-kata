package _arrays

import (
	"sort"
)

/*
Insert Interval

You are given an array of non-overlapping intervals.
intervals where intervals[i] = [start_i, end_i] represents the start and the end time of the
ith interval. intervals is initially sorted in ascending order by start_i.
You are given another interval newInterval = [start, end].
Insert newInterval into intervals such that intervals is still sorted
in ascending order by start_i and also intervals still does not have any overlapping intervals.
You may merge the overlapping intervals if needed. Return intervals after adding newInterval.

Note: Intervals are non-overlapping if they have no common point.
For example, [1,2] and [3,4] are non-overlapping, but [1,2] and [2,3] are overlapping.

Example 1: Input: intervals = [[1,3],[4,6]], newInterval = [2,5] Output: [[1,6]]
Explanation: [2,5] overlaps with [1,3] and [4,6], so all three are merged into [1,6].

Example 2: Input: intervals = [[1,2],[3,5],[9,10]], newInterval = [6,7] Output: [[1,2],[3,5],[6,7],[9,10]]
Explanation: [6,7] does not overlap with any existing interval, so it is simply inserted between [3,5] and [9,10].
*/
func insert(intervals [][]int, newInterval []int) [][]int {
	var rslt [][]int
	n := len(intervals)
	i := 0

	// we are lower than the new interval
	for i < n && intervals[i][1] < newInterval[0] {
		rslt = append(rslt, intervals[i])
		i++
	}

	// adding new interval
	for i < n && intervals[i][0] <= newInterval[1] {
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}

		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
		i++
	}
	rslt = append(rslt, newInterval)

	// appending rest of the intervals slice
	for i < n {
		rslt = append(rslt, intervals[i])
		i++
	}

	return rslt
}

/*
Merge Interval

Given an array of intervals where intervals[i] = [start_i, end_i], merge all overlapping
intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

You may return the answer in any order.

Note: Intervals are non-overlapping if they have no common point.
For example, [1, 2] and [3, 4] are non-overlapping, but [1, 2] and [2, 3] are overlapping.

Example 1:
Input: intervals = [[1,3],[1,5],[6,7]]
Output: [[1,5],[6,7]]

Example 2:
Input: intervals = [[1,2],[2,3]]
Output: [[1,3]]
*/
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	rslt := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := intervals[len(rslt)-1]
		curr := intervals[i]

		if curr[0] <= last[1] {
			if curr[1] > last[1] {
				last[1] = curr[1]
			}
		} else {
			rslt = append(rslt, curr)
		}
	}

	return rslt
}

/*
Non-overlapping Intervals
Given an array of  intervals where intervals[i] = [start_i, end_i], return the
minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.

Note: Intervals are non-overlapping even if they have a common point.
For example, [1, 3] and [2, 4] are overlapping, but [1, 2] and [2, 3] are non-overlapping.

Example 1:
Input: intervals = [[1,2],[2,4],[1,4]]
Output: 1

Explanation: After [1,4] is removed, the rest of the intervals are non-overlapping.

Example 2:
Input: intervals = [[1,2],[2,4]]
Output: 0
*/
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	var overlaps int
	lastEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		curr := intervals[i]
		if curr[0] < lastEnd {
			overlaps++
		} else {
			lastEnd = curr[1]
		}

	}

	return overlaps
}

type Interval struct {
	start, end int
}

/*
Meeting rooms

Given an array of meeting time interval objects consisting of start and end times
[[start_1,end_1],[start_2,end_2],...] (start_i < end_i), determine if a person
could add all meetings to their schedule without any conflicts.

Note: (0,8),(8,10) is not considered a conflict at 8

Input: intervals = [(0,30),(5,10),(15,20)]
Output: false

Input: intervals = [(5,8),(9,15)]
Output: true
*/
func canAttendMeetings(intervals []Interval) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].start < intervals[j].start
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i-1].end > intervals[i].start {
			return false
		}
	}

	return true
}
