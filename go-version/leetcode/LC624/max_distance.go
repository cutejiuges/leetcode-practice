package lc624

import "math"

func maxDistance(arrays [][]int) int {
	mn, mx := math.MaxInt/2, math.MinInt/2
	ans := 0
	for _, arr := range arrays {
		first, last := arr[0], arr[len(arr)-1]
		ans = max(ans, mx-first, last-mn)
		mn = min(mn, first)
		mx = max(mx, last)
	}
	return ans
}
