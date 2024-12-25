package LC3218

import "sort"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/25 下午10:45
 * @FilePath: minimum_cost
 * @Description: 贪心策略
 */

func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int {
	//将切割成本存储在一个列表中，第一个表示成本，第二个表示切割方向，0表示横切，1表示竖切
	cuts := make([][2]int, 0)

	for i := range horizontalCut {
		cuts = append(cuts, [2]int{horizontalCut[i], 0})
	}
	for j := range verticalCut {
		cuts = append(cuts, [2]int{verticalCut[j], 1})
	}
	sort.Slice(cuts, func(i, j int) bool {
		return cuts[i][0] > cuts[j][0]
	})

	horizontalCnt, verticalCnt := 1, 1 //初始化横块和竖块的数量
	ans := 0
	for i := range cuts {
		cost := cuts[i][0]
		if cuts[i][1] == 0 { //横向切割
			ans += verticalCnt * cost //当横切的时候，用一块的成本乘以竖块的数量作为本次切割成本
			horizontalCnt++           //切完了之后，横快的数量会变多
		} else { //竖向切割
			ans += horizontalCnt * cost //当竖切的时候，用一块的成本乘以横块的数量作为本次切割成本
			verticalCnt++               //切万了之后，竖块的数量会增加
		}
	}
	return ans
}
