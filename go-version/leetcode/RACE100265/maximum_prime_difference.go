package RACE100265

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/14 上午10:49
 * @FilePath: maximum_prime_difference
 * @Description: 素数的最大距离，艾氏筛法
 */
const MAX = 101

var Prime [MAX + 1]bool

func init() {
	for i := 0; i <= MAX; i++ {
		Prime[i] = true
	}
	Prime[0], Prime[1] = false, false
	for i := 2; i < MAX; i++ {
		if Prime[i] {
			for j := i * 2; j <= MAX; j += i {
				Prime[j] = false
			}
		}
	}
}

func maximumPrimeDifference(nums []int) int {
	length := len(nums)
	low, high := 0, length-1
	for low <= high {
		if !Prime[nums[low]] {
			low++
		}
		if !Prime[nums[high]] {
			high--
		}
		if Prime[nums[low]] && Prime[nums[high]] {
			break
		}
	}
	return high - low
}
