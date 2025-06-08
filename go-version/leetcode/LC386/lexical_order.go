package LC386

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/6/8 上午8:40
 * @FilePath: lexical_order
 * @Description:
 */

func lexicalOrder(n int) []int {
	var ans []int
	num := 1
	for i := 0; i < n; i++ {
		ans = append(ans, num)
		if num*10 <= n {
			num *= 10
		} else {
			for num%10 == 9 || num >= n {
				num /= 10
			}
			num++
		}
	}
	return ans
}
