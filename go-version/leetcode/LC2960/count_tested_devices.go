package LC2960

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/10 下午11:37
 * @FilePath: count_tested_devices
 * @Description: 统计已测试设备，模拟遍历即可
 */

// 1. 初始化 ans=0，表示需要减一的次数。
// 2. 对于电池i，其实际百分比即batteryPercentages[i]-ans，若此值>0，那么此后每个电池都要-1,直接将ans+1即可
// 3. 答案就是ans
func countTestedDevices(batteryPercentages []int) int {
	ans := 0
	for i := range batteryPercentages {
		if batteryPercentages[i] > ans {
			ans++
		}
	}
	return ans
}
