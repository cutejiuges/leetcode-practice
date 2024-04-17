/*
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024-04-01 08:10:38
 * @LastEditors: cutejiuge cutejiuge@163.com
 * @LastEditTime: 2024-04-01 08:33:19
 * @FilePath: /leetcode-practice/leetcode/LC2810/final_string.go
 * @Description: 故障键盘，直接模拟即可
 */
package lc2810

// 直接模拟，遇到i直接进行反转，其他字母顺序添加
func finalString(s string) string {
	bytes := make([]rune, 0)
	for _, ch := range s {
		if ch == 'i' {
			reverse(bytes)
		} else {
			bytes = append(bytes, ch)
		}
	}
	return string(bytes)
}

func reverse(bytes []rune) {
	for i, length := 0, len(bytes); i < length/2; i++ {
		bytes[i], bytes[length-i-1] = bytes[length-i-1], bytes[i]
	}
}

// 使用标记位和双端队列进行模拟
func finalString2(s string) string {
	needReverse, ans := false, []rune{}
	for _, ch := range s {
		if ch == 'i' {
			needReverse = !needReverse
		} else {
			if !needReverse {
				ans = append(ans, ch)
			} else {
				ans = append([]rune{ch}, ans...)
			}
		}
	}
	if needReverse {
		reverse(ans)
	}
	return string(ans)
}
