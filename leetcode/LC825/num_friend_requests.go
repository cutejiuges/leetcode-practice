package LC825

import "sort"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/17 下午1:52
 * @FilePath: num_friend_requests
 * @Description:
 */

func numFriendRequests(ages []int) int {
	sort.Ints(ages)
	left, right := 0, 0
	ans := 0
	for _, age := range ages {
		// 条件1的筛选
		if age <= 14 {
			continue
		}

		// 筛选条件2
		for ages[left]*2 <= age+14 {
			left++
		}
		for right+1 < len(ages) && ages[right+1] <= age {
			right++
		}
		ans += right - left
	}
	return ans
}
