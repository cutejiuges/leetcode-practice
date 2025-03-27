package LC2712

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/27 下午11:37
 * @FilePath: minimum_cost
 * @Description: 脑筋急转弯
 */

func minimumCost(s string) int64 {
	n := len(s)
	var ans int64
	for i := 1; i < n; i++ {
		if s[i] != s[i-1] {
			ans += int64(min(i, n-i))
		}
	}
	return ans
}
