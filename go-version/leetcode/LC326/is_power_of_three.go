package LC326

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/13 下午11:21
 * @FilePath: is_power_of_three
 * @Description: 递归判断
 */

func isPowerOfThree(n int) bool {
	if n < 3 {
		return n == 1
	}
	return n%3 == 0 && isPowerOfThree(n/3)
}
