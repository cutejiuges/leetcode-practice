package LC3206

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/26 下午11:15
 * @FilePath: number_of_alternating_groups
 * @Description:
 */

func numberOfAlternatingGroups(colors []int) int {
	n := len(colors)
	ans := 0
	for i := 0; i < n; i++ {
		if idx := i + n; colors[idx%n] != colors[(idx-1)%n] && colors[idx%n] != colors[(idx+1)%n] {
			ans++
		}
	}
	return ans
}
