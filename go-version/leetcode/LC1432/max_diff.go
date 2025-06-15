package LC1432

import (
	"strconv"
	"strings"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/6/15 下午8:14
 * @FilePath: max_diff
 * @Description: 字典序替换
 */

func maxDiff(num int) int {
	str := strconv.Itoa(num)
	mx, mn := num, num
	for _, ch := range str {
		if ch != '9' {
			mx = convert(str, string(ch), "9")
			break
		}
	}

	if str[0] != '1' {
		mn = convert(str, str[:1], "1")
	} else {
		for _, ch := range str[1:] {
			if ch >= '2' {
				mn = convert(str, string(ch), "0")
				break
			}
		}
	}

	return mx - mn
}

func convert(source, old, new string) int {
	str := strings.ReplaceAll(source, old, new)
	num, _ := strconv.Atoi(str)
	return num
}
