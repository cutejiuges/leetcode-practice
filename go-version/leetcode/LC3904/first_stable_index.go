package LC3904

func firstStableIndex(nums []int, k int) int {
	length := len(nums)
	// 构造最小后缀数组
	suffixMinArr := make([]int, length)
	suffixMinArr[length-1] = nums[length-1]
	for i := length - 2; i >= 0; i-- {
		suffixMinArr[i] = min(suffixMinArr[i+1], nums[i])
	}

	// 顺序前缀遍历
	prefixMaxVal := nums[0]
	for i := 0; i < length; i++ {
		prefixMaxVal = max(nums[i], prefixMaxVal)
		if prefixMaxVal-suffixMinArr[i] <= k {
			return i
		}
	}
	return -1
}
