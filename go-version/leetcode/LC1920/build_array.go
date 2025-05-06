package LC1920

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/6 下午10:47
 * @FilePath: build_array
 * @Description: 直接模拟
 */

func buildArray(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = nums[nums[i]]
	}
	return ans
}
