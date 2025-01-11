package race416_q2

import (
	"math"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/9/22 上午10:45
 * @FilePath: min_number_of_seconds
 * @Description: 注意二分的边界 => https://blog.csdn.net/Sx_WWj/article/details/129092427
 */

// 求出某个workerTime在time时间内，所能挖走的最大高度
func calHeight(workerTime, time int) int {
	//高度h满足 h*(h+1) / 2 * workTime <= time
	//通过二分来确定
	low, high := 0, int(math.Sqrt(float64(time*2)/float64(workerTime)))+1
	for low < high {
		mid := low + (high-low+1)/2
		condition := mid * (mid + 1) / 2 * workerTime
		if condition <= time {
			low = mid
		} else {
			high = mid - 1
		}
	}
	return low
}

// 根据最大的workerTime确定耗时的上界
func maxWorkTime(height int, workerTimes []int) int64 {
	mx := workerTimes[0]
	for _, time := range workerTimes {
		if mx < time {
			mx = time
		}
	}
	return int64(height * (height + 1) / 2 * mx)
}

func canFinish(height, time int, workerTime []int) bool {
	h := 0
	for _, wkTime := range workerTime {
		h += calHeight(wkTime, time)
		if h >= height {
			return true
		}
	}
	return false
}

func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	low, high := int64(0), maxWorkTime(mountainHeight, workerTimes)
	for low < high {
		mid := low + (high-low)/2
		if canFinish(mountainHeight, int(mid), workerTimes) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
