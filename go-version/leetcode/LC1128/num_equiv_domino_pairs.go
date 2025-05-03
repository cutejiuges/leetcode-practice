package LC1128

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/4 上午12:41
 * @FilePath: num_equiv_domino_pairs
 * @Description: 枚举存储即可
 */

func numEquivDominoPairs(dominoes [][]int) int {
	visited := make([][]int, 10)
	for i := range visited {
		visited[i] = make([]int, 10)
	}
	var ans int
	for _, domino := range dominoes {
		x, y := min(domino[0], domino[1]), max(domino[0], domino[1])
		ans += visited[x][y]
		visited[x][y]++
	}
	return ans
}
