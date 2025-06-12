/*
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024-03-24 12:39:58
 * @LastEditors: cutejiuge cutejiuge@163.com
 * @LastEditTime: 2024-03-24 12:46:05
 * @FilePath: /leetcode-practice/math_util/calculate.go
 * @Description: 数学工具类，实现常见的数学运算
 */
package math_util

func Max[T int | float32 | float64](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func Min[T int | float32 | float64](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

func Abs[T int | float32 | float64](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
