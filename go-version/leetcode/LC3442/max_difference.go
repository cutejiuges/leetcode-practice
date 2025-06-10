package LC3442

import "math"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/6/10 下午10:55
 * @FilePath: max_difference
 * @Description: 直接遍历，按值进行更新即可
 */

func maxDifference(s string) int {
	cnt := make([]int, 26)
	for _, ch := range s {
		cnt[ch-'a']++
	}
	mx, mn := math.MinInt, math.MaxInt
	for _, num := range cnt {
		if num == 0 {
			continue
		}
		if num&1 == 1 {
			mx = max(mx, num)
		} else {
			mn = min(mn, num)
		}
	}
	return mx - mn
}
