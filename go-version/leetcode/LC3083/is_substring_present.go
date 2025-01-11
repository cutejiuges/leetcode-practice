package LC3083

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/26 下午11:07
 * @FilePath: is_substring_present
 * @Description: 枚举状态
 */

func isSubstringPresent(s string) bool {
	exits := [26][26]bool{}
	for i := 1; i < len(s); i++ {
		a, b := s[i-1]-'a', s[i]-'a'
		exits[a][b] = true
		if exits[b][a] {
			return true
		}
	}
	return false
}
