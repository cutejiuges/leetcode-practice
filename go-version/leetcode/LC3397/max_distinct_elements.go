package LC3397

import (
	"math"
	"sort"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/10/18 下午10:27
 * @FilePath: max_distinct_elements
 * @Description: 排序后进行贪心的比对
 */

func maxDistinctElements(nums []int, k int) int {
	var ans int
	pre := math.MinInt
	sort.Ints(nums)
	for _, num := range nums {
		v := min(max(pre+1, num-k), num+k)
		if v > pre {
			pre = v
			ans++
		}
	}
	return ans
}
