package LC3261

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/13 下午9:43
 * @FilePath: count_k_constraint_substrings
 * @Description:
 */

func countKConstraintSubstrings(s string, k int, queries [][]int) []int64 {
	ans := make([]int64, len(queries))
	for i, qry := range queries {
		ans[i] = helper(s[qry[0]:qry[1]+1], k)
	}
	return ans
}

func helper(s string, k int) int64 {
	n, ans := len(s), int64(0)
	cnt := [2]int{0, 0}

	for left, right := 0, 0; right < n; right++ {
		cnt[s[right]-'0']++
		for ; cnt[0] > k && cnt[1] > k; left++ {
			cnt[s[left]-'0']--
		}
		ans += int64(right - left + 1)
	}
	return ans
}
