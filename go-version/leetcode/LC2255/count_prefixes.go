package LC2255

import "strings"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/24 下午10:33
 * @FilePath: count_prefixes
 * @Description: 直接使用库函数
 */

func countPrefixes(words []string, s string) int {
	var ans int
	for _, word := range words {
		if strings.HasPrefix(s, word) {
			ans++
		}
	}
	return ans
}
