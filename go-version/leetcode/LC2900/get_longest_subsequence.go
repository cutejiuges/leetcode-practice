package LC2900

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/15 下午9:55
 * @FilePath: get_longest_subsequence
 * @Description:
 */

func getLongestSubsequence(words []string, groups []int) []string {
	n := len(groups)
	ans := make([]string, 0)
	for i := 0; i < n; i++ {
		if i == n-1 || groups[i] != groups[i+1] {
			ans = append(ans, words[i])
		}
	}
	return ans
}
