package leetcode

import "sort"

// https://leetcode-cn.com/problems/non-overlapping-intervals
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(a, b int) bool {
		if intervals[a][0] != intervals[b][0] {
			return intervals[a][0] < intervals[b][0]
		} else {
			return intervals[a][1] < intervals[b][1]
		}
	})
	count := 0
	for i, j := 0, 1; i < len(intervals) && j < len(intervals); {
		if intervals[i][1] > intervals[j][0] {
			count++
			if intervals[i][1] < intervals[j][1] {
				i = i - 1
			} else {
				i = j - 1
			}
		} else {
			i = j - 1
		}
		i++
		j++
	}
	return count
}
