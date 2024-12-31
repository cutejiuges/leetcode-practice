package LC3219

import "sort"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/31 下午8:07
 * @FilePath: minimum_cost
 * @Description: 还是贪心，复用3218的方法进行尝试
 */

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int64 {
	//按照横切和竖切的代价进行排序，放在二维数组里面，0表示横切，1表示竖切
	var cuts [][2]int
	for i := range horizontalCut {
		cuts = append(cuts, [2]int{horizontalCut[i], 0})
	}
	for i := range verticalCut {
		cuts = append(cuts, [2]int{verticalCut[i], 1})
	}
	sort.Slice(cuts, func(i, j int) bool {
		return cuts[i][0] > cuts[j][0]
	})

	horizontalNum, verticalNum := 1, 1
	ans := int64(0)
	for _, cut := range cuts {
		if cut[1] == 0 {
			//如果是横切，那么横切代价*竖块数量，同时横块数量增加
			ans += int64(cut[0] * verticalNum)
			horizontalNum++
		} else {
			//如果是竖切，那么竖切代价*横块数量，同时竖块数量增加
			ans += int64(cut[0] * horizontalNum)
			verticalNum++
		}
	}
	return ans
}
