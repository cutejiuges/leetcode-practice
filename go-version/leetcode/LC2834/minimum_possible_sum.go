/*
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024-03-09 13:24:14
 * @LastEditors: cutejiuge cutejiuge@163.com
 * @LastEditTime: 2024-03-09 15:24:46
 * @FilePath: /leetcode-practice/leetcode/LC2834/minimum_possible_sum.go
 * @Description: lc2834 日题练习，考虑使用贪心的方式来解决
 */
package lc2834

/**
 * @description: 题面详情见lc2384,为了让所得和最小，我们从1开始递增相加，因为任意2数之和不能为target，因此在比target的数字中，一旦存在i
 * 则不能存在target - i，因此比target小的数字序列以target/2为界限，之前的数字加入后，之后的数字不能放入。
 * 若target/2已经满足了n个数字的要求，直接计算前n项和；如果没有满足，则从target开始，在递增相加直到n即可。
 * @param {int} n，数字的数量要求，要求数组长度为n时的最小和
 * @param {int} target，限制条件，数组中不能存在2个数的和是target
 * @return {*} 计算出来的最小和
 */
func minimumPossibleSum(n int, target int) int {
	mod := int(1e9 + 7)
	bound := target / 2
	if n <= bound {
		return ((1 + n) * n / 2) % mod
	}
	return (((1+bound)*bound/2)%mod + ((target+target+(n-bound)-1)*(n-bound)/2)%mod) % mod
}
