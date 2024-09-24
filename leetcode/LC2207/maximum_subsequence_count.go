package LC2207

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/9/24 下午11:49
 * @FilePath: maximum_subsequence_count
 * @Description:
 */

func maximumSubsequenceCount(text string, pattern string) int64 {
	cnt0, cnt1 := int64(0), int64(0)
	res := int64(0)
	for _, ch := range text {
		if pattern[1] == byte(ch) {
			cnt1++
			res += cnt0
		}
		if pattern[0] == byte(ch) {
			cnt0++
		}
	}
	res += max(cnt0, cnt1)
	return res
}
