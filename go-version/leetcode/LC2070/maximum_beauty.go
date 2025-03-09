package LC2070

import "sort"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/9 下午3:58
 * @FilePath: maximum_beauty
 * @Description: 排序+二分查找
 */

func maximumBeauty(items [][]int, queries []int) []int {
	//先按照item的price进行排序
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})
	//更换前缀最大值
	n := len(items)
	for i := 0; i < n-1; i++ {
		items[i+1][1] = max(items[i+1][1], items[i][1])
	}

	//按照二分进行处理查询
	n = len(queries)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = query(items, queries[i])
	}
	return ans
}

func query(items [][]int, q int) int {
	low, high := 0, len(items)
	for low < high {
		mid := low + ((high - low) >> 1)
		if items[mid][0] > q {
			high = mid
		} else {
			low = mid + 1
		}
	}
	if low == 0 {
		return 0
	}
	return items[low-1][1]
}
