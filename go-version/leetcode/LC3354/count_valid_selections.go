package LC3354

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/10/28 下午11:28
 * @FilePath: count_valid_selections
 * @Description: 前缀和进行分解判断即可
 */

func countValidSelections(nums []int) int {
	var sum, cur, ans int
	for _, num := range nums {
		sum += num
	}
	for _, num := range nums {
		cur += num
		if num == 0 && cur*2 == sum {
			ans += 2
		}
		if num == 0 && abs(cur*2-sum) == 1 {
			ans += 1
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
