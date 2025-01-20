package lc2239

import "math"

func findClosestNumber(nums []int) int {
	ans := math.MaxInt
	for _, num := range nums {
		if abs(num) < abs(ans) {
			ans = num
		} else if abs(num) == abs(ans) {
			ans = max(ans, num)
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
