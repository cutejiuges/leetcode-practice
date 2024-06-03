package LC1103

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/3 下午11:00
 * @FilePath: distribute_candies
 * @Description:
 */

func distributeCandies(candies int, numPeople int) []int {
	ans := make([]int, numPeople)
	for i := 0; candies > 0; {
		ans[i%numPeople] += min(candies, i+1)
		candies -= min(candies, i+1)
		i++
	}
	return ans
}
