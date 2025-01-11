package LC3143

import "sort"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/8/3 下午9:31
 * @FilePath: max_points_inside_square
 * @Description:
 */

/*按照切比雪夫距离转移到第一象限，按距离排序*/
func maxPointsInsideSquare(points [][]int, s string) int {
	graph := make(map[int][]int)
	for i, point := range points {
		key := max(abs(point[0]), abs(point[1]))
		graph[key] = append(graph[key], i)
	}

	visited := make([]bool, 26)
	keys := []int{}
	ans := 0
	for key := range graph {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys {
		idx := graph[key]
		for _, xi := range idx {
			j := s[xi] - 'a'
			if visited[j] {
				return ans
			}
			visited[j] = true
		}
		ans += len(idx)
	}
	return ans
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
