package LC3306

import "strings"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/13 下午11:30
 * @FilePath: count_of_substrings
 * @Description: 可以将恰好型问题转换成2个至少型问题
 */

func countOfSubstrings(word string, k int) int64 {
	return calculate(word, k) - calculate(word, k+1)
}

func calculate(ss string, k int) int64 {
	var ans int64
	vowelCntMap := make(map[byte]int)
	consonantCnt := 0
	low := 0
	for i := 0; i < len(ss); i++ {
		if strings.ContainsRune("aeiou", rune(ss[i])) {
			vowelCntMap[ss[i]]++
		} else {
			consonantCnt++
		}

		leftChar := ss[low]
		for len(vowelCntMap) >= 5 && consonantCnt >= k {
			if vowelCntMap[leftChar] > 0 {
				vowelCntMap[leftChar]--
				if vowelCntMap[leftChar] <= 0 {
					delete(vowelCntMap, leftChar)
				}
			} else {
				consonantCnt--
			}
			low++
		}
		ans += int64(low)
	}
	return ans
}
