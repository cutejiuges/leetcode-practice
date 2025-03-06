package LC2588

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/6 下午11:35
 * @FilePath: beautiful_subarrays
 * @Description:
 */

/*
根据题意可以知，由于每次操作中需要从子数组中选择两个不同的数分别减去 2^k，
使得子数组中所有元素均变为 0，由此可知对于子数组中所有元素 2^k 出现的次数之和必须是偶数。
换一种说法，即对于二进制中第 i 位，则数组中所有元素第 i 位为 1 的数目一定为偶数，
则此时满足数组中所有元素第 i 位的异或和一定为 0。
*/

func beautifulSubarrays(nums []int) int64 {
	mp := map[int]int{0: 1}
	mask, ans := 0, int64(0)
	for _, num := range nums {
		mask ^= num
		ans += int64(mp[mask])
		mp[mask]++
	}
	return ans
}
