package LC2141

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/12/1 上午9:12
 * @FilePath: max_run_time
 * @Description:
 */

func maxRunTime(n int, batteries []int) int64 {
	var total int64
	for _, num := range batteries {
		total += int64(num)
	}

	low, high := int64(0), total/int64(n)+1
	for low+1 < high {
		mid := low + (high-low)/2
		var temp int64
		for _, num := range batteries {
			temp += min(mid, int64(num))
		}
		if int64(n)*mid <= temp {
			low = mid
		} else {
			high = mid
		}
	}
	return low
}
