package LC869

import "strconv"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/10 下午10:52
 * @FilePath: reordered_power_of_2
 * @Description: 回溯
 */

var visited []bool

func reorderedPowerOf2(n int) bool {
	chars := []byte(strconv.Itoa(n))
	visited = make([]bool, len(chars))
	return backtrack(chars, 0, 0)
}

func backtrack(chars []byte, idx, num int) bool {
	if idx == len(chars) {
		return powOf2(num)
	}
	for i := 0; i < len(chars); i++ {
		if num == 0 && chars[i] == '0' {
			continue
		}
		if visited[i] {
			continue
		}
		visited[i] = true
		if backtrack(chars, idx+1, num*10+int(chars[i]-'0')) {
			return true
		}
		visited[i] = false
	}
	return false
}

func powOf2(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}
