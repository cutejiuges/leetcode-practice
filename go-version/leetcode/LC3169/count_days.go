package LC3169

import "sort"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/7/11 上午9:04
 * @FilePath: count_days
 * @Description: 按照左端点升序排序，合并区间
 */

func countDays(days int, meetings [][]int) int {
	ans := days
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	start, end := meetings[0][0], meetings[0][1]
	for i := 1; i < len(meetings); i++ {
		if meetings[i][0] > end { //区间不相交
			ans -= end - start + 1
			start = meetings[i][0]
		}
		end = max(end, meetings[i][1])
	}
	ans -= end - start + 1
	return ans
}
