package lc2270

func waysToSplitArray(nums []int) int {
	left, right := int64(0), int64(0)
	for _, num := range nums {
		right += int64(num)
	}

	ans, n := 0, len(nums)
	for i := 0; i < n-1; i++ {
		left += int64(nums[i])
		right -= int64(nums[i])
		if left >= right {
			ans++
		}
	}
	return ans
}
