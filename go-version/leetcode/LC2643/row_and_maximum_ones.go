package LC2643

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/22 下午1:01
 * @FilePath: row_and_maximum_ones
 * @Description: 遍历计数即可
 */

func rowAndMaximumOnes(mat [][]int) []int {
	maxCntOfOnes, index := -1, -1
	for i := 0; i < len(mat); i++ {
		onesCnt := 0
		for _, m := range mat[i] {
			onesCnt += m
		}
		if onesCnt > maxCntOfOnes {
			maxCntOfOnes = onesCnt
			index = i
		}
	}
	return []int{index, maxCntOfOnes}
}
