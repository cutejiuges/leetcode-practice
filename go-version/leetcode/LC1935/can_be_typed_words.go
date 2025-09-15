package LC1935

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/9/15 下午11:35
 * @FilePath: can_be_typed_words
 * @Description: 字符串模拟，使用二进制压缩减少空间消耗
 */

func canBeTypedWords(text, brokenLetters string) int {
	var mask int
	for _, ch := range brokenLetters {
		mask |= 1 << (ch - 'a')
	}
	ans, ok := 0, 1
	for _, ch := range text {
		if ch == ' ' {
			ans += ok
			ok = 1
			continue
		}
		if mask&(1<<(ch-'a')) != 0 {
			ok = 0
		}
	}
	ans += ok
	return ans
}
