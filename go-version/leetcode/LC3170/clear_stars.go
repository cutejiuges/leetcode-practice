package LC3170

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/6/7 下午11:45
 * @FilePath: clear_stars
 * @Description: 使用栈进行模拟
 */

func clearStars(s string) string {
	n := len(s)
	chars := []byte(s)
	lists := make([][]int, 26)
	for i := 0; i < 26; i++ {
		lists[i] = make([]int, 0)
	}

	for i := 0; i < n; i++ {
		if chars[i] != '*' {
			lists[chars[i]-'a'] = append(lists[chars[i]-'a'], i)
		} else {
			for j, list := range lists {
				if len(list) > 0 {
					idx := list[len(list)-1]
					lists[j] = list[:len(list)-1]
					chars[idx] = '*'
					break
				}
			}
		}
	}

	var ans string
	for _, ch := range chars {
		if ch != '*' {
			ans += string(ch)
		}
	}
	return ans
}
