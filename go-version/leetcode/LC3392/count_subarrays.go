package LC3392

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/27 上午8:28
 * @FilePath: count_subarrays
 * @Description: 遍历模拟即可
 */

func countSubarrays(nums []int) int {
	var ans int
	n := len(nums)
	for i := 1; i < n-1; i++ {
		if nums[i] == (nums[i-1]+nums[i+1])*2 {
			ans++
		}
	}
	return ans
}
