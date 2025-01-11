/*
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024-03-27 22:36:30
 * @LastEditors: cutejiuge cutejiuge@163.com
 * @LastEditTime: 2024-03-27 23:23:47
 * @FilePath: /leetcode-practice/leetcode/LC2580/count_ways.go
 * @Description: 合并区间的问题，考虑排序+遍历方式实现
 */
package lc2580

import "sort"

/**
* @description: 考虑使用区间合并的方式来实现，将所有有交集的区间进行合并，假设合并后的结果有x个，那么每个集合都可以放进第一组或者第二组，组合方式就能有2^x个
* @param {[][]int} ranges 待分组的区间你
* @return {*} 分组的组合数
 */
func countWays(ranges [][]int) int {
	mod := int(1e9 + 7)
	//先按照区间的起点升序排列
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i][0] == ranges[j][0] {
			return ranges[i][1] < ranges[j][1]
		}
		return ranges[i][0] < ranges[j][0]
	})

	cnt := 1
	end := ranges[0][1]
	for _, rg := range ranges {
		if end < rg[0] {
			cnt++
		}
		if rg[1] > end {
			end = rg[1]
		}
	}
	return quickPow(2, cnt, mod)
}

func quickPow(base, pow, mod int) int {
	ans := 1
	for pow != 0 {
		if pow&1 == 1 {
			ans = ans * base % mod
		}
		base = base * base % mod
		pow >>= 1
	}
	return ans
}
