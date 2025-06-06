package LC2434

import "math"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/6/6 下午10:25
 * @FilePath: robot_with_string
 * @Description:
 */

func robotWithString(s string) string {
	n := len(s)
	sufMin := make([]byte, n+1)
	sufMin[n] = math.MaxUint8
	for i := n - 1; i >= 0; i-- {
		sufMin[i] = min(sufMin[i+1], s[i])
	}

	ans := make([]byte, 0, n)
	stack := sufMin[:0]
	for i := 0; i < n; i++ {
		stack = append(stack, s[i])
		for len(stack) > 0 && stack[len(stack)-1] <= sufMin[i+1] {
			ans = append(ans, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}
	return string(ans)
}
