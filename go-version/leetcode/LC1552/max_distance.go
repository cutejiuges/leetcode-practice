package lc1552

import "sort"

func maxDistance(position []int, m int) int {
	sort.Ints(position)
	low, high := 1, position[len(position)-1]
	for low < high {
		mid := (low + high + 1) >> 1
		if check(position, m, mid) {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return low
}

func check(position []int, m, mid int) bool {
	pre, cnt := position[0], 1
	for _, cur := range position {
		if cur-pre >= mid {
			cnt++
			pre = cur
		}
	}
	return cnt >= m
}
