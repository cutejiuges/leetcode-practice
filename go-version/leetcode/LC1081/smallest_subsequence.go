package LC1081

func smallestSubsequence(s string) string {
	cnt, visited := make([]int, 26), make([]bool, 26)
	// 统计每一个字母的个数
	for _, ch := range s {
		cnt[ch-'a']++
	}
	var ans []byte
	for _, ch := range []byte(s) {
		cnt[ch-'a']--
		if visited[ch-'a'] {
			continue
		}
		// 如果结果的顶字符比当前字符更大且后续的待选集中还有顶字符，就把顶字符吐出来
		for len(ans) > 0 && ch < ans[len(ans)-1] && cnt[ans[len(ans)-1]-'a'] > 0 {
			visited[ans[len(ans)-1]-'a'] = false
			ans = ans[:len(ans)-1]
		}
		ans = append(ans, ch)
		visited[ch-'a'] = true
	}
	return string(ans)
}
