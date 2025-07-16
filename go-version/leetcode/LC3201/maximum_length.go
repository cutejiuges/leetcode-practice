package LC3201

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/7/16 下午11:41
 * @FilePath: maximum_length
 * @Description:
 */

func maximumLength(nums []int) int {
	var ans int
	patterns := [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	for _, pattern := range patterns {
		cnt := 0
		for _, num := range nums {
			if num&1 == pattern[cnt&1] {
				cnt++
			}
		}
		ans = max(ans, cnt)
	}
	return ans
}
