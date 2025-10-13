package LC2273

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/10/13 下午10:54
 * @FilePath: remove_anagrams
 * @Description: 顺序遍历进行判断
 */

func removeAnagrams(words []string) []string {
	var ans []string
	if len(words) <= 0 {
		return ans
	}
	ans = append(ans, words[0])
	idx := 0
	for i := 1; i < len(words); i++ {
		if anagram(ans[idx], words[i]) {
			continue
		}
		ans = append(ans, words[i])
		idx++
	}
	return ans
}

func anagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	cnt := make([]byte, 26)
	for _, v := range a {
		cnt[v-'a']++
	}
	for _, v := range b {
		cnt[v-'a']--
	}
	for _, v := range cnt {
		if v != 0 {
			return false
		}
	}
	return true
}
