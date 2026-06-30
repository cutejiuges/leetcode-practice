package LC1358

func numberOfSubstrings(s string) int {
	var ans, left int
	cnt := make([]int, 3)
	for _, ch := range s {
		cnt[ch-'a']++
		for cnt[0] > 0 && cnt[1] > 0 && cnt[2] > 0 {
			cnt[s[left]-'a']--
			left++
		}
		ans += left
	}
	return ans
}
