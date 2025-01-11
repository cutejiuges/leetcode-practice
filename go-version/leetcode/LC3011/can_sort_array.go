package LC3011

import "math/bits"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/7/13 上午9:52
 * @FilePath: can_sort_array
 * @Description:
 */

func canSortArray(nums []int) bool {
	var (
		lastCnt      = 0 //上一个元素二进制位1的个数
		lastGroupMax = 0 //上一个分组的最大值
		curGroupMax  = 0 //当前分组的最大值
	)

	for _, num := range nums {
		curCnt := bits.OnesCount(uint(num)) //记录当前元素的二进制位1个数
		if lastCnt != curCnt {              //如果二进制位1个数不同，需要新开一个分组
			lastCnt = curCnt           //将元素记录切换到新的分组
			lastGroupMax = curGroupMax //最大值也放到新分组
			curGroupMax = num          //当前分组第一个元素就是当前分组的最大值
		} else { //如果二进制位1个数相等，在同一个分组
			curGroupMax = max(curGroupMax, num) //更新最大值即可
		}
		if num < lastGroupMax { //如果当前元素小于上一个分组的最大值，则不可能升序了
			return false
		}
	}
	return true
}
