package LC2476

import (
	datastruct "cutejiuge/leetcode-practice/data_struct"
	"sort"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/7/10 上午12:03
 * @FilePath: closest_node
 * @Description: 二叉搜索树可以遍历查询，或者直接放进数组里面二分查找
 */

func closestNodes(root *datastruct.TreeNode, queries []int) [][]int {
	sortedSlice := datastruct.MiddleOrder(root)
	res := make([][]int, len(queries))
	for i, v := range queries {
		res[i] = processQuery(sortedSlice, v)
	}
	return res
}

func processQuery(array []int, target int) []int {
	maxVal, minVal := -1, -1
	idx := sort.SearchInts(array, target)
	if idx < len(array) {
		maxVal = array[idx]
		if array[idx] == target {
			minVal = array[idx]
			return []int{minVal, maxVal}
		}
	}
	if idx != 0 {
		minVal = array[idx-1]
	}
	return []int{minVal, maxVal}
}
