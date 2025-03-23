package LC2116

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/23 上午11:30
 * @FilePath: can_be_valid
 * @Description: 维护未匹配的左括号的最值
 */

func canBeValid(s string, locked string) bool {
	n := len(s)
	if n&1 == 1 {
		return false
	}

	mn, mx := 0, 0
	for i := 0; i < n; i++ {
		if locked[i] == '1' { //不能更改当前字符
			var cur int
			if s[i] == '(' {
				cur = 1
			} else {
				cur = -1
			}
			mx += cur
			if mx < 0 { //当前前缀字符串中右括号比左括号多，不能匹配了
				return false
			}
			mn += cur
		} else { //当前字符可以改
			mx++ //改成左括号，未匹配数+1
			mn-- //改成右括号，未匹配数-1
		}
		mn = max(0, mn)
	}
	return mn == 0
}
