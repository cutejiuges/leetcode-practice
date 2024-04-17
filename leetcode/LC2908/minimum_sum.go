/*
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024-03-29 09:26:46
 * @LastEditors: cutejiuge cutejiuge@163.com
 * @LastEditTime: 2024-03-31 10:09:42
 * @FilePath: /leetcode-practice/leetcode/LC2908/minimum_sum.go
 * @Description: 山形数组问题，可以考虑暴力或者前缀后缀和
 */
package lc2908

import (
	"cutejiuge/leetcode-practice/math_util"
	"math"
)

func minimumSum(nums []int) int {
	length, res := len(nums), math.MaxInt
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			for k := j + 1; k < length; k++ {
				if nums[i] < nums[j] && nums[k] < nums[j] {
					res = math_util.Min(res, nums[i]+nums[j]+nums[k])
				}
			}
		}
	}
	if res < 1000 {
		return res
	}
	return -1
}
