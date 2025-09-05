package LC2749

import "math/bits"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/9/5 下午11:43
 * @FilePath: make_the_integer_zero
 * @Description: 枚举位数，进行转换
 */

func makeTheIntegerZero(num1, num2 int) int {
	// num1 = k*2^i + k*num2, 则k*2^i = num1 - k*num2 = x
	// 问题转换为是否存在k满足x = k*2^i
	for k := 1; ; k++ {
		x := num1 - k*num2
		if x < k {
			return -1
		}
		if k >= bits.OnesCount(uint(x)) {
			return k
		}
	}
}
