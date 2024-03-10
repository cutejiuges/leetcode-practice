/*
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024-03-11 00:06:24
 * @LastEditors: cutejiuge cutejiuge@163.com
 * @LastEditTime: 2024-03-11 00:22:32
 * @FilePath: /leetcode-practice/leetcode/LC2129/capitalize_title.go
 * @Description: lc2129,将每个长度大于2的单词首字母变成大写，其他字母均变成小写，直接遍历
 */
package lc2129

func capitalizeTitle(title string) string {
	n := len(title)
	left, right := 0, 0 //记录单词的区间，左闭右开
	title += " "        //在单词最末价格空格，统一边界处理
	bytes := []byte(title)
	for right < n {
		for bytes[right] != ' ' {
			right++
		}
		//如果单词长度比2大，首字母i大写，其他全部保持小写
		if right-left > 2 {
			bytes[left] = toUpper(bytes[left])
			left++
		}
		for left < right {
			bytes[left] = toLower(bytes[left])
			left++
		}
		left = right + 1
		right++
	}
	bytes = bytes[:n]
	return string(bytes)
}

func toUpper(ch byte) byte {
	if ch >= 'a' && ch <= 'z' {
		ch -= 32
	}
	return ch
}

func toLower(ch byte) byte {
	if ch >= 'A' && ch <= 'Z' {
		ch += 32
	}
	return ch
}
