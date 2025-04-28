package LC2302

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/28 下午11:29
 * @FilePath: count_subarrays
 * @Description: 使用滑动窗口算法进行计算，越短越合法类型
 */

func countSubarrays(nums []int, k int64) int64 {
	var ans, score int64
	var left int
	n := len(nums)

	for i := 0; i < n; i++ {
		score += int64(nums[i])
		for score*int64(i-left+1) >= k {
			score -= int64(nums[left])
			left++
		}
		ans += int64(i - left + 1)
	}
	return ans
}
