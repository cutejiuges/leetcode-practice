/*
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024-03-25 09:12:07
 * @LastEditors: cutejiuge cutejiuge@163.com
 * @LastEditTime: 2024-03-25 09:16:51
 * @FilePath: /leetcode-practice/leetcode/LC70/climb_stairs.go
 * @Description: 爬楼梯，经典动态规划问题
 */
package lc70

func climbStairs(n int) int {
	pre, res := 1, 1
	for i := 2; i <= n; i++ {
		pre, res = res, pre+res
	}
	return res
}
