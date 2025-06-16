package LC2016

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/6/16 上午8:15
 * @FilePath: maximum_difference
 * @Description: 前缀最小值
 */

func maximumDifference(nums []int) int {
	ans, preMin := -1, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] <= preMin {
			preMin = nums[i]
		} else {
			ans = max(ans, nums[i]-preMin)
		}
	}
	return ans
}
