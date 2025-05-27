package LC2894

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/27 下午11:22
 * @FilePath: difference_of_sums
 * @Description: 直接遍历就行了
 */

func differenceOfSums(n, m int) int {
	var ans int
	for i := 1; i <= n; i++ {
		if i%m != 0 {
			ans += i
		} else {
			ans -= i
		}
	}
	return ans
}
