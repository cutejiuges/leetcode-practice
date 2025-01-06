package LC2274

import "sort"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/1/6 上午8:16
 * @FilePath: max_consecutive
 * @Description: 直接排序
 */

func maxConsecutive(bottom, top int, special []int) int {
	special = append(special, bottom-1, top+1)
	sort.Ints(special)
	n := len(special)
	ans := 0
	for i := 1; i < n; i++ {
		ans = max(ans, special[i]-special[i-1]-1)
	}
	return ans
}
