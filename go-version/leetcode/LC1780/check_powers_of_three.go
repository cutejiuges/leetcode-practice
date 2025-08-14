package LC1780

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/14 下午10:59
 * @FilePath: check_powers_of_three
 * @Description: 转换3进制
 */

func checkPowersOfThree(num int) bool {
	return powerSumOfX(num, 3)
}

func powerSumOfX(num, n int) bool {
	if num < 1 {
		return false
	}
	for num != 0 {
		if num%n > 1 {
			return false
		}
		num /= n
	}
	return true
}
