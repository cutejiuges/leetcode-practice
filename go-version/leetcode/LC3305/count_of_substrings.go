package LC3305

import "strings"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/9/30 下午9:42
 * @FilePath: count_of_substrings
 * @Description:
 */

// 非常经典的滑动窗口问题，计算刚好为k个的，可以先计算至少为k的和至少为k+1的，然后直接相减
func countOfSubstrings(word string, k int) int64 {
	return calcLeastK(word, k) - calcLeastK(word, k+1)
}

func calcLeastK(ss string, k int) int64 {
	cnt1 := make(map[byte]int) //元音字母各自的个数
	cnt2 := 0                  //辅音字母的个数
	low := 0
	ret := int64(0)
	for _, ch := range ss {
		//统计元音和辅音字母的个数
		if strings.ContainsRune("aeiou", ch) {
			cnt1[byte(ch)]++
		} else {
			cnt2++
		}

		//如果满足条件的话从左侧开始缩小范围
		for len(cnt1) >= 5 && cnt2 >= k {
			pop := ss[low]
			if strings.ContainsRune("aeiou", rune(pop)) {
				cnt1[pop]--
				if cnt1[pop] == 0 {
					delete(cnt1, pop)
				}
			} else {
				cnt2--
			}
			low++
		}
		//当前最小串及其左边所有字符拼接都可以组成满足条件的子串
		ret += int64(low)
	}
	return ret
}
