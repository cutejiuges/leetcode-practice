package LC1925

import "math"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/12/8 上午8:23
 * @FilePath: count_triples
 * @Description: 枚举遍历
 */

func countTriples(n int) int {
	var ans int
	for a := 1; a < n; a++ {
		for b := 1; b < a && a*a+b*b <= n*n; b++ {
			c := a*a + b*b
			rt := int(math.Sqrt(float64(c)))
			if rt*rt == c {
				ans++
			}
		}
	}
	return ans * 2
}
