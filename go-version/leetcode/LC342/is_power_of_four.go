package LC342

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/15 下午9:14
 * @FilePath: is_power_of_four
 * @Description:
 */

func isPowerOfFour(n int) bool {
	if n < 1 {
		return false
	}
	for n%4 == 0 {
		n /= 4
	}
	return n == 1
}
