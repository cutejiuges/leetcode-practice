package LC3356

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/22 上午12:13
 * @FilePath: min_zero_array
 * @Description:
 */

func minZeroArray(nums []int, queries [][]int) int {
	queryLength := len(queries)
	low, high := 0, queryLength+1
	for low < high {
		mid := low + (high-low)/2
		if check(nums, queries, mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	if low <= queryLength {
		return low
	}
	return -1
}

func check(nums []int, queries [][]int, k int) bool {
	n := len(nums)
	diffArray := make([]int, n)
	for _, query := range queries[:k] {
		diffArray[query[0]] += query[2]
		if query[1] < n-1 {
			diffArray[query[1]+1] -= query[2]
		}
	}

	sumDiff := 0
	for i, diff := range diffArray {
		sumDiff += diff
		if nums[i] > sumDiff {
			return false
		}
	}
	return true
}
