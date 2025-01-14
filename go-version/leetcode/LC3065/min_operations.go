package lc3065

func minOperations(nums []int, k int) int {
	ans := 0
	for i := range nums {
		if nums[i] < k {
			ans++
		}
	}
	return ans
}
