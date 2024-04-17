package lc2007

import "sort"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/18 00:43
 * @FilePath: find_origin_array
 * @Description: 从倍数数组还原原数组
 */

func findOriginalArray(changed []int) []int {
	sort.Ints(changed)
	cnt := make(map[int]int)

	//统计没一个数字出现的频率
	for _, num := range changed {
		cnt[num]++
	}

	//再次遍历数组，根据数字频率找答案
	ans := []int{}
	for _, num := range changed {
		//如果数字频次已经是0,说明这是一个二倍数，已经被处理，跳过处理下一个
		if cnt[num] == 0 {
			continue
		}
		//原数字频次减小
		cnt[num]--

		//判断当前数字的二倍数是否存在，如果存在，将其频次减小，标记为是一个二倍数
		//如果不存在，说明某数字没有二倍数，则直接返回一个空数组
		if cnt[num*2] == 0 {
			return []int{}
		}
		cnt[num*2]--
		ans = append(ans, num)
	}
	return ans
}
