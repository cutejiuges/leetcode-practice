package LC1550

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/11 下午8:27
 * @FilePath: three_consecutive_odds
 * @Description: 直接一次遍历计算即可
 */

func threeConsecutiveOdds(nums []int) bool {
	for i := 0; i < len(nums)-2; i++ {
		if nums[i]&1 == 1 && nums[i+1]&1 == 1 && nums[i+2]&1 == 1 {
			return true
		}
	}
	return false
}
