package LC2009

import "sort"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/8 下午11:28
 * @FilePath: min_operations
 * @Description: 一个思路，正难则反
 */

func minOperations(nums []int) int {
	length := len(nums)
	//统计数字，进行去重操作
	cntMap := make(map[int]bool)
	for _, num := range nums {
		cntMap[num] = true
	}
	//存放去重结果
	distinctSortNums := []int{}
	for num, _ := range cntMap {
		distinctSortNums = append(distinctSortNums, num)
	}
	//对去重结果进行排序
	sort.Ints(distinctSortNums)

	ans, idx := length, 0
	for i, low := range distinctSortNums {
		high := low + length - 1
		for idx < len(distinctSortNums) && distinctSortNums[idx] <= high {
			ans = min(ans, length-(idx-i+1))
			idx++
		}
	}
	return ans
}
