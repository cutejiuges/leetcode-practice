package LC1304

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/9/7 下午10:22
 * @FilePath: sum_zero
 * @Description: 按照相反数依次填入即可
 */

func sumZero(n int) []int {
	ans := make([]int, n)
	low, high := 0, n-1
	k := n
	for low < high {
		ans[low] = -k
		ans[high] = k
		k--
		low++
		high--
	}
	if n&1 == 1 {
		ans[low] = 0
	}
	return ans
}
