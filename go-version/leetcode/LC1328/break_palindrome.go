package LC1328

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/5 下午11:27
 * @FilePath: break_palindrome
 * @Description: 破坏回文串
 */

func breakPalindrome(palindrome string) string {
	n := len(palindrome)
	if n <= 1 {
		return ""
	}

	ss := []byte(palindrome)
	for i := 0; i < n; i++ {
		if ss[i] != 'a' {
			ss[i] = 'a'
			return string(ss)
		}
	}
	ss[n-1] = 'b'
	return string(ss)
}
