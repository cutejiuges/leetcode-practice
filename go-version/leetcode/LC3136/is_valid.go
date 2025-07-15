package LC3136

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/7/15 下午11:56
 * @FilePath: is_valid
 * @Description:
 */

func isValid(word string) bool {
	if len(word) < 3 {
		return false
	}
	var containsVowel, containsConsonant bool
	for _, ch := range word {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
				ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
				containsVowel = true
			} else {
				containsConsonant = true
			}
		} else if !(ch >= '0' && ch <= '9') {
			return false
		}
	}
	return containsVowel && containsConsonant
}
