package LC66

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2026/1/1 下午10:18
 * @FilePath: plus_one
 * @Description:
 */

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		digits[i] = (digits[i] + 1) % 10
		if digits[i] != 0 {
			return digits
		}
	}
	ans := make([]int, n+1)
	ans[0] = 1
	return ans
}
