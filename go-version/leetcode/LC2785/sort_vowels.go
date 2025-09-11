package LC2785

import "slices"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/9/11 下午11:29
 * @FilePath: sort_vowels
 * @Description:
 */

func sortVowels(s string) string {
	vowelSet := map[byte]struct{}{
		'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}, 'A': {}, 'E': {}, 'I': {}, 'O': {}, 'U': {},
	}
	vowels := make([]byte, 0)
	ss := []byte(s)
	for _, ch := range ss {
		if _, ok := vowelSet[ch]; ok {
			vowels = append(vowels, ch)
		}
	}
	slices.Sort(vowels)
	idx := 0
	for i, ch := range ss {
		if _, ok := vowelSet[ch]; ok {
			ss[i] = vowels[idx]
			idx++
		}
	}
	return string(ss)
}
