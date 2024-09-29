package LC2073

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/9/29 下午11:23
 * @FilePath: time_required_to_buy
 * @Description: 模拟，提取公因数
 */

// 数组循环模拟
func timeRequiredToBuy2(tickets []int, k int) int {
	res, idx := 0, 0
	length := len(tickets)
	for tickets[k] > 0 {
		if tickets[idx] > 0 {
			tickets[idx]--
			res++
		}
		idx = (idx + 1) % length
	}
	return res
}

// 一次遍历计算
func timeRequiredToBuy(tickets []int, k int) int {
	res := 0
	for i, t := range tickets {
		if i <= k {
			res += min(t, tickets[k])
		} else {
			res += min(t, tickets[k]-1)
		}
	}
	return res
}
