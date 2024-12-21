package LC2545

import "sort"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/21 下午8:39
 * @FilePath: sort_the_students
 * @Description:
 */

func sortTheStudents(score [][]int, k int) [][]int {
	sort.Slice(score, func(i, j int) bool {
		return score[i][k] > score[j][k]
	})
	return score
}
