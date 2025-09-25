package LC165

import "strings"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/9/23 下午11:30
 * @FilePath: compare_version
 * @Description: 按点好拆解之后，补位比较大小即可
 */

func compareVersion(version1, version2 string) int {
	v1, v2 := strings.Split(version1, "."), strings.Split(version2, ".")
	len1, len2 := len(v1), len(v2)
	for i := 0; i < len1 || i < len2; i++ {
		var a, b int
		if i < len1 {
			a = parseInt(v1[i])
		}
		if i < len2 {
			b = parseInt(v2[i])
		}
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
	}
	return 0
}

func parseInt(s string) int {
	ans, leadZero := 0, true
	for i := 0; i < len(s); i++ {
		if s[i] == '0' && leadZero {
			continue
		} else if s[i] != '0' {
			leadZero = false
		}
		ans *= 10
		ans += int(s[i]) - int('0')
	}
	return ans
}
