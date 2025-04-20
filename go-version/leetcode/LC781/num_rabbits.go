package LC781

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/20 上午8:09
 * @FilePath: num_rabbits
 * @Description: 分组计数
 */

func numRabbits(answers []int) int {
	var ans int
	mp := make(map[int]int)
	for _, num := range answers {
		if mp[num] == 0 {
			ans += num + 1
			mp[num] = num
		} else {
			mp[num]--
		}
	}
	return ans
}
