package LC1493

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/24 上午7:09
 * @FilePath: longest_subarray
 * @Description: 转换题意求只有1个0的最长子串，滑动窗口
 */

func longestSubarray(nums []int) int {
	var cnt0, ans, left int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			cnt0++
		}
		for cnt0 > 1 {
			if nums[left] == 0 {
				cnt0--
			}
			left++
		}
		ans = max(ans, i-left)
	}
	return ans
}
