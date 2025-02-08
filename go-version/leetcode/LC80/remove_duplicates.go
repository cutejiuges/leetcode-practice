package lc80

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	curSize, maxSize := 2, len(nums)
	for i := 2; i < maxSize; i++ {
		if nums[i] != nums[curSize-2] {
			nums[curSize] = nums[i]
			curSize++
		}
	}
	return min(curSize, maxSize)
}
