package LC3194

import (
	"math"
	"sort"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/16 上午12:43
 * @FilePath: minimum_average
 * @Description:
 */

func minimumAverage(nums []int) float64 {
	sort.Ints(nums)
	n := len(nums)
	avg := math.MaxFloat64
	for i := 0; i < n/2; i++ {
		avg = min(avg, float64(nums[i]+nums[n-i-1])/2)
	}
	return avg
}
