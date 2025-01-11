package LC2748

import "cutejiuge/leetcode-practice/math_util"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/20 下午10:47
 * @FilePath: count_beautiful_pairs
 * @Description: 暴力遍历 + 哈希计数
 */

func countBeautifulPairs(nums []int) (ans int) {
	//return violentCount(nums)
	return hashCount(nums)
}

func violentCount(nums []int) (ans int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 10 {
			nums[i] /= 10
		}
		for j := i + 1; j < n; j++ {
			if math_util.Gcd(nums[i], nums[j]%10) == 1 {
				ans++
			}
		}
	}
	return ans
}

func hashCount(nums []int) (ans int) {
	cnt := make([]int, 10)
	for _, num := range nums {
		for d := 1; d < 10; d++ {
			if math_util.Gcd(d, num%10) == 1 {
				ans += cnt[d]
			}
		}
		for num > 10 {
			num /= 10
		}
		cnt[num]++
	}
	return ans
}
