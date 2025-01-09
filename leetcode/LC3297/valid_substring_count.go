package LC3297

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/1/9 下午11:12
 * @FilePath: valid_substring_count
 * @Description:
 */

func validSubstringCount(word1 string, word2 string) int64 {
	if len(word1) < len(word2) {
		return 0
	}

	cnt := [26]int{}
	remain := 0
	for _, ch := range word2 {
		cnt[ch-'a']++
		if cnt[ch-'a'] == 1 {
			remain++
		}
	}

	window := [26]int{}
	left := 0
	ans := int64(0)
	for _, ch := range word1 {
		i := int(ch - 'a')
		window[i]++
		if window[i] == cnt[i] {
			remain--
		}
		for remain == 0 {
			i = int(word1[left] - 'a')
			if window[i] == cnt[i] {
				remain++
			}
			window[i]--
			left++
		}
		ans += int64(left)
	}
	return ans
}
