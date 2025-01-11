package LC575

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/2 下午8:43
 * @FilePath: distribute_candies
 * @Description:
 */

func distributeCandies(candyType []int) int {
	ans := len(candyType) / 2
	set := map[int]struct{}{}
	for _, v := range candyType {
		set[v] = struct{}{}
	}
	ans = min(ans, len(set))
	return ans
}
