package LC3184

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/22 下午11:07
 * @FilePath: count_complete_day_pairs
 * @Description: 使用两数之和的思路来实现
 */

func countCompleteDayPairs(hours []int) int {
	mp := make(map[int]int)
	ans := 0
	for _, hour := range hours {
		ans += mp[(24-hour%24)%24]
		mp[hour%24]++
	}
	return ans
}
