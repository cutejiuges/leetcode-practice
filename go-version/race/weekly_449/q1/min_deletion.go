package q1

import "sort"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/11 上午10:39
 * @FilePath: min_deletion
 * @Description:
 */

func minDeletion(s string, k int) int {
	cnt := make([]int, 26)
	for _, ch := range s {
		cnt[ch-'a']++
	}
	sort.Ints(cnt)
	var ans int
	for i := 0; i < 26-k; i++ {
		ans += cnt[i]
	}
	return ans
}
