package LC1578

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/11/3 下午11:22
 * @FilePath: min_cost
 * @Description: 贪心 + 栈
 */

func minCost(colors string, neededTime []int) int {
	preColor := colors[0]
	preTime, cost := neededTime[0], 0
	for i := 1; i < len(colors); i++ {
		if preColor == colors[i] {
			cost += min(preTime, neededTime[i])
			preTime = max(preTime, neededTime[i])
		} else {
			preColor = colors[i]
			preTime = neededTime[i]
		}
	}
	return cost
}
