package LC520

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/23 下午4:20
 * @FilePath: detect_capital_use
 * @Description: 遍历检查即可
 */

func detectCapitalUse(s string) bool {
	if len(s) == 0 {
		return true
	}
	if s[0] >= 'A' && s[0] <= 'Z' {
		return allUppercase(s[1:]) || allLowercase(s[1:])
	}
	return allLowercase(s[1:])
}

func allUppercase(s string) bool {
	for _, ch := range s {
		if !(ch >= 'A' && ch <= 'Z') {
			return false
		}
	}
	return true
}

func allLowercase(s string) bool {
	for _, ch := range s {
		if !(ch >= 'a' && ch <= 'z') {
			return false
		}
	}
	return true
}
