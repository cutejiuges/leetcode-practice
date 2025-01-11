package LC3242

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/9 下午4:37
 * @FilePath: neighbor_sum
 * @Description:
 */

type NeighborSum struct {
	res [][2]int
}

func Constructor(grid [][]int) NeighborSum {
	m, n := len(grid), len(grid[0])
	res := make([][2]int, m*n)

	dir := [8][2]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
		{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
	}

	for i, row := range grid {
		for j, val := range row {
			for k, d := range dir {
				dx, dy := i+d[0], j+d[1]
				if dx >= 0 && dx < m && dy >= 0 && dy < n {
					res[val][k/4] += grid[dx][dy]
				}
			}
		}
	}

	return NeighborSum{
		res: res,
	}
}

func (this *NeighborSum) AdjacentSum(value int) int {
	return this.res[value][0]
}

func (this *NeighborSum) DiagonalSum(value int) int {
	return this.res[value][1]
}
