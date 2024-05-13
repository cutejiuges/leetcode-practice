package LC994

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/13 下午10:36
 * @FilePath: oranges_rotting
 * @Description: 腐烂的橘子，可以使用广度优点搜索方式来实现
 */

var dir = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func orangesRotting(grid [][]int) int {
	//如果是空矩阵，直接返回0
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rowNum, colNum := len(grid), len(grid[0])
	queue := make([][]int, 0)
	healthyOrangeNum := 0
	ans := 0

	//一次遍历，记录好橘子的数量和烂橘子的位置
	for i := 0; i < rowNum*colNum; i++ {
		row, col := i/colNum, i%colNum
		if grid[row][col] == 1 {
			healthyOrangeNum++
		}
		if grid[row][col] == 2 {
			queue = append(queue, []int{row, col})
		}
	}

	//只要有好橘子，且存在传染源，执行传染的动作
	for healthyOrangeNum > 0 && len(queue) > 0 {
		ans++
		n := len(queue)
		for i := 0; i < n; i++ {
			//当前元素出队
			pos := queue[0]
			queue = queue[1:]
			//遍历相邻的4个方向
			for j := 0; j < 4; j++ {
				newRow, newCol := pos[0]+dir[j][0], pos[1]+dir[j][1]
				//在grid范围内，且是个好橘子，则将好橘子腐烂，并加入传染源的位置
				if newRow >= 0 && newRow < rowNum && newCol >= 0 && newCol < colNum && grid[newRow][newCol] == 1 {
					grid[newRow][newCol] = 2
					healthyOrangeNum--
					queue = append(queue, []int{newRow, newCol})
				}
			}
		}
	}
	if healthyOrangeNum > 0 {
		return -1
	}
	return ans
}
