package LC440

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/6/9 下午11:37
 * @FilePath: find_kth_number
 * @Description:
 */

func findKthNumber(n, k int) int {
	cur := 1
	k--
	for k > 0 {
		step := getSteps(cur, int64(n))
		if step <= k {
			k -= step
			cur++
		} else {
			k--
			cur *= 10
		}
	}
	return cur
}

func getSteps(cur int, n int64) int {
	step := 0
	first, last := int64(cur), int64(cur)
	for first <= n {
		step += int((min(last, n) - first) + 1)
		first *= 10
		last = last*10 + 9
	}
	return step
}
