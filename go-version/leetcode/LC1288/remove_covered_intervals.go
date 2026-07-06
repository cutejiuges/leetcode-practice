package LC1288

import "sort"

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] > intervals[j][1]
	})

	ans, right := 1, intervals[0][1]
	length := len(intervals)
	for i := 1; i < length; i++ {
		if intervals[i][1] > right {
			ans++
			right = intervals[i][1]
		}
	}
	return ans
}
