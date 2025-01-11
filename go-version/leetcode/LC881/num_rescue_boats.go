package LC881

import "sort"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/10 下午11:51
 * @FilePath: num_rescue_boats
 * @Description:
 */

func numRescueBoats(people []int, limit int) int {
	ans := 0
	sort.Ints(people)
	low, high := 0, len(people)-1
	for low <= high {
		if people[low]+people[high] > limit {
			high--
		} else {
			high--
			low++
		}
		ans++
	}
	return ans
}
