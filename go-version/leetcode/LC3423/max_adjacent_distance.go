package LC3423

import "cutejiuge/leetcode-practice/math_util"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/6/12 下午11:12
 * @FilePath: max_adjacent_distance
 * @Description: 遍历即可
 */

func maxAdjacentDistance(nums []int) int {
	n := len(nums)
	ans := math_util.Abs(nums[0] - nums[n-1])
	for i := 1; i < n; i++ {
		ans = max(ans, math_util.Abs(nums[i]-nums[i-1]))
	}
	return ans
}
