/*
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024-03-08 00:03:03
 * @LastEditors: cutejiuge cutejiuge@163.com
 * @LastEditTime: 2024-03-09 13:11:51
 * @FilePath: /leetcode-practice/leetcode/LC2575/divisibility_array.go
 * @Description: 迭代还原字符串中的数字，需要注意的是字符串很长可能会溢出，所以需要将用数学的同余原理处理一下
 */
package lc2575

func divisibilityArray(word string, m int) []int {
	num := 0
	res := make([]int, 0)
	for _, ch := range word {
		num = (num*10 + int((ch - '0'))) % m
		if num%m == 0 {
			res = append(res, 1)
		} else {
			res = append(res, 0)
		}
	}
	return res
}
