package LC2131

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/25 下午6:52
 * @FilePath: longest_palindrome
 * @Description:
 */

func longestPalindrome(words []string) int {
	mp := make(map[string]int)
	for _, word := range words {
		mp[word]++
	}
	st := make(map[string]struct{})
	ans := 0
	odd := false

	for word, cnt := range mp {
		reversedWord := string(word[1]) + string(word[0])
		if reversedWord == word {
			if cnt&1 == 1 {
				odd = true
			}
			ans += (cnt / 2) * 2 * 2
		} else if _, ok := st[reversedWord]; !ok {
			st[word] = struct{}{}
			st[reversedWord] = struct{}{}
			ans += min(mp[word], mp[reversedWord]) * 2 * 2
		}
	}
	if odd {
		ans += 2
	}
	return ans
}
