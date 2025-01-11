package LC908

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/20 上午9:08
 * @FilePath: smallest_range_1
 * @Description:
 */

func smallestRangeI(nums []int, k int) int {
	mn, mx := nums[0], nums[0]
	n := len(nums)

	for i := 1; i < n; i++ {
		if nums[i] < mn {
			mn = nums[i]
		}
		if nums[i] > mx {
			mx = nums[i]
		}
	}

	if mx-mn <= k*2 {
		return 0
	}
	return mx - mn - k*2
}
