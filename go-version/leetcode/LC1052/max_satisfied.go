package LC1052

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/23 上午8:51
 * @FilePath: max_satisfied
 * @Description: 滑动窗口思路解决
 */

// 可以先将本来就不生气的客户数量直接找出来，同时将本来就不生气的这些时刻的客户数量记为0,标记这些客户不可变更
// 再次遍历数组，维护一个长度为minutes的滑动窗口，统计滑动窗口内的可变用户，把这些用户的最大数量找出来，加到不生气的用户上面，即为答案
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	n, ans := len(customers), 0
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			ans += customers[i]
			customers[i] = 0
		}
	}

	cur, maxu := 0, 0
	for i := 0; i < n; i++ {
		cur += customers[i]
		if i >= minutes {
			cur -= customers[i-minutes]
		}
		maxu = max(maxu, cur)
	}
	return maxu + ans
}
