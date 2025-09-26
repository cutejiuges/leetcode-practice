package LC611

import "sort"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/9/26 下午9:15
 * @FilePath: triangle_number
 * @Description: 排序 + 双指针
 */

func triangleNumber(nums []int) int {
	n, ans := len(nums), 0
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		k := i
		for j := i + 1; j < n; j++ {
			for k+1 < n && nums[k+1] < nums[i]+nums[j] {
				k++
			}
			ans += max(k-j, 0)
		}
	}
	return ans
}
