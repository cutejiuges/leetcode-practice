package LC1399

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/23 下午10:01
 * @FilePath: count_largest_group
 * @Description: 遍历分组计数
 */

func countLargestGroup(n int) int {
	mp := make(map[int][]int)
	ans, maxSize := 0, 0
	for i := 1; i <= n; i++ {
		sum := getDigitSum(i)
		mp[sum] = append(mp[sum], i)
		if len(mp[sum]) > maxSize {
			maxSize = len(mp[sum])
			ans = 1
		} else if len(mp[sum]) == maxSize {
			ans++
		}
	}
	return ans
}

func getDigitSum(x int) int {
	ans, t := 0, x
	for t > 0 {
		ans += t % 10
		t /= 10
	}
	return ans
}
