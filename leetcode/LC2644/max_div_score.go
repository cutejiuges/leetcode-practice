package LC2644

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/22 上午3:17
 * @FilePath: max_div_score
 * @Description: 遍历+哈希记录
 */

func maxDivScore(nums, divisors []int) int {
	score, ans := -1, -1

	for _, divisor := range divisors {
		curScore := 0
		for _, num := range nums {
			if num%divisor == 0 {
				curScore++
			}
		}
		if curScore > score || (curScore == score && divisor < ans) {
			score = curScore
			ans = divisor
		}
	}
	return ans
}
