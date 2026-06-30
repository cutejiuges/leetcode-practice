package LC1846

import "sort"

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	sort.Ints(arr)
	length := len(arr)
	arr[0] = 1
	for i := 1; i < length; i++ {
		arr[i] = min(arr[i-1]+1, arr[i])
	}
	return arr[length-1]
}
