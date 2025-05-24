package LC2942

import "strings"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/24 下午4:49
 * @FilePath: find_words_contatining
 * @Description:
 */

func findWordsContaining(words []string, x byte) []int {
	ans := make([]int, 0)
	for i, word := range words {
		if strings.Contains(word, string(x)) {
			ans = append(ans, i)
		}
	}
	return ans
}
