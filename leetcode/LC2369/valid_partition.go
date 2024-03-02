/*
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024-03-01 22:57:58
 * @LastEditors: cutejiuge cutejiuge@163.com
 * @LastEditTime: 2024-03-01 23:16:08
 * @FilePath: /leetcode-practice/LC2369/valid_partition.go
 * @Description: leetcode 2369题，考虑动态规划来解决
 */
package lc2369

func judgeTwoNum(elems []int) bool {
	return elems[0] == elems[1]
}

func judgeThreeNum(elems []int) bool {
	if elems[0] == elems[1] && elems[1] == elems[2] {
		return true
	}
	if elems[0]+1 == elems[1] && elems[1]+1 == elems[2] {
		return true
	}
	return false
}

func validPartition(nums []int) bool {
	length := len(nums)
	dp := make([]bool, length+1)
	dp[0] = true
	for i := 1; i <= length; i++ {
		if i >= 2 {
			dp[i] = dp[i-2] && judgeTwoNum(nums[i-2:i])
		}
		if i >= 3 {
			dp[i] = dp[i] || (dp[i-3] && judgeThreeNum(nums[i-3:i]))
		}
	}
	return dp[length]
}
