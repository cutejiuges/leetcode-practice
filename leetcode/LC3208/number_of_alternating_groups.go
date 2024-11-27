package LC3208

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/27 下午8:26
 * @FilePath: number_of_alternating_groups
 * @Description: 直接拼接一个双倍数组进行统计
 */

func numberOfAlternatingGroups(colors []int, k int) int {
	n := len(colors)
	ans, cnt := 0, 0
	for i := 0; i < 2*n; i++ {
		if i > 0 && colors[i%n] == colors[(i-1)%n] {
			cnt = 0
		}
		cnt++
		if i >= n && cnt >= k {
			ans++
		}
	}
	return ans
}
