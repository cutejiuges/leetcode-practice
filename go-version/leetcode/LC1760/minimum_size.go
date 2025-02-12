package lc1760

func minimumSize(nums []int, maxOperations int) int {
	mx := 0
	for _, num := range nums {
		mx = max(mx, num)
	}

	low, high := 0, mx
	for low+1 < high {
		mid := (low + high) >> 1
		if check(nums, maxOperations, mid) {
			high = mid
		} else {
			low = mid
		}
	}
	return high
}

func check(nums []int, maxOperations, mid int) bool {
	var cnt int
	for _, num := range nums {
		cnt += (num - 1) / mid
	}
	return cnt <= maxOperations
}
