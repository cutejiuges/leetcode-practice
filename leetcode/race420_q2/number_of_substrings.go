package race420_q2

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/20 上午11:24
 * @FilePath: number_of_substrings
 * @Description:
 */

func numberOfSubstrings(s string, k int) int {
	n, ans := len(s), 0
	for i := 0; i < n; i++ {
		cnt := make(map[byte]int)
		mx := 0
		for j := i; j < n; j++ {
			cnt[s[j]]++
			mx = max(mx, cnt[s[j]])
			if mx >= k {
				ans++
			}
		}
	}
	return ans
}
