package LC633

import "math"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/4 上午9:11
 * @FilePath: judge_square_sum
 * @Description: 平方数之和
 */

func judgeSquareSum(c int) bool {
	//遍历枚举
	for a := 0; a*a <= (c+1)/2; a++ {
		b := math.Sqrt(float64(c - a*a))
		if b == math.Floor(b) {
			return true
		}
	}
	return false
}
