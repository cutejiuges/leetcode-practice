package LC2843

import "strconv"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/11 下午9:54
 * @FilePath: count_symmetric_integers
 * @Description: 遍历数字进行判断
 */

func countSymmetricIntegers(low, high int) int {
	var ans int
	for i := low; i <= high; i++ {
		str := strconv.Itoa(i)
		n := len(str)
		if n&1 == 1 {
			continue
		}
		judge := 0
		for j := 0; j < n/2; j++ {
			judge += int(str[j])
		}
		for j := n / 2; j < n; j++ {
			judge -= int(str[j])
		}
		if judge == 0 {
			ans++
		}
	}
	return ans
}
