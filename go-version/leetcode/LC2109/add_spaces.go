package LC2109

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/30 上午8:41
 * @FilePath: add_spaces
 * @Description: 双指针
 */

func addSpaces(s string, spaces []int) string {
	ans := make([]byte, 0)
	spaceIdx := 0
	for i, ch := range s {
		if spaceIdx < len(spaces) && i == spaces[spaceIdx] {
			ans = append(ans, ' ')
			spaceIdx++
		}
		ans = append(ans, byte(ch))
	}
	return string(ans)
}
