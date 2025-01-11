package LC2663

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/22 上午2:50
 * @FilePath: smallest_beautiful_string
 * @Description: 贪心构造
 */

func smallestBeautifulString(s string, k int) string {
	for i := len(s) - 1; i >= 0; i-- {
		blockedChar := make(map[byte]bool)
		for j := 1; j < 3; j++ {
			if i-j >= 0 {
				blockedChar[s[i-j]] = true
			}
		}

		for j := 1; j < 4; j++ {
			if int(s[i]-'a')+j+1 <= k && !blockedChar[s[i]+byte(j)] {
				return generate(s, i, j)
			}
		}
	}
	return ""
}

func generate(s string, idx, offset int) string {
	res := []byte(s)
	res[idx] += byte(offset)
	for i := idx + 1; i < len(s); i++ {
		blockedChar := make(map[byte]bool)
		for j := 1; j < 3; j++ {
			if i-j >= 0 {
				blockedChar[res[i-j]] = true
			}
		}
		for c := byte('a'); c <= byte('c'); c++ {
			if !blockedChar[c] {
				res[i] = c
				break
			}
		}
	}
	return string(res)
}
