package LC2566

import (
	"strconv"
	"strings"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/6/14 下午8:56
 * @FilePath: min_max_difference
 * @Description:
 */

func minMaxDifference(num int) int {
	str := strconv.Itoa(num)
	mx, mn := num, num
	for _, ch := range str {
		if ch != '9' {
			mx, _ = strconv.Atoi(strings.ReplaceAll(str, string(ch), "9"))
			break
		}
	}
	mn, _ = strconv.Atoi(strings.ReplaceAll(str, str[:1], "0"))
	return mx - mn
}
