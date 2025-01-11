/*
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024-03-06 23:15:20
 * @LastEditors: cutejiuge cutejiuge@163.com
 * @LastEditTime: 2024-03-06 23:26:46
 * @FilePath: /leetcode-practice/leetcode/LC2917/find_kor.go
 * @Description: 遍历数位，如果数组元素中超过k个元素此数位的bit是1,将res的此数位置为1即可
 */
package lc2917

func findKOr(nums []int, k int) int {
	res := 0
	for i := 0; i < 31; i++ {
		cnt := 0
		for _, num := range nums {
			if num>>i&1 == 1 {
				cnt++
			}
		}
		if cnt >= k {
			res = res | 1<<i
		}
	}
	return res
}
