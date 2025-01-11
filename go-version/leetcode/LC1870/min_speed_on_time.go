package LC1870

import (
	"math"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/2 下午1:05
 * @FilePath: min_speed_on_time
 * @Description:
 */

func minSpeedOnTime(dist []int, hour float64) int {
	//为了避免计算误差，先将hour*100转换为整数计算
	hr := int(math.Round(hour * 100))
	n := len(dist)
	//假定速度足够快，每一段路都能在1小时内完成，时间至少应该大于路程段数-1
	if hr <= 100*(n-1) {
		return -1
	}

	//如果能到的话考虑二分
	low, high := 1, int(1e7)
	for low < high {
		mid := low + (high-low)>>1
		if judge(dist, hr, mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func judge(dist []int, hr int, speed int) bool {
	t, n := 0, len(dist)
	// 前面的n-1段路程中，每一段贡献的时间需要向上取整
	for i := 0; i < n-1; i++ {
		t += (dist[i]-1)/speed + 1
	}
	//这段时间内理论上跑过的路程为
	t *= speed
	//加上最后一段路程
	t += dist[n-1]

	//在hr时间内跑过的路程为 hr*speed，以此为判断依据
	return hr*speed >= 100*t
}
