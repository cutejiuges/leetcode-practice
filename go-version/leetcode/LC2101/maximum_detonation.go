package LC2101

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/7/22 下午11:25
 * @FilePath: maximum_detonation
 * @Description: 建图之后广搜
 */

func canBomb(bombs [][]int, x, y int) bool {
	dx := bombs[x][0] - bombs[y][0]
	dy := bombs[x][1] - bombs[y][1]
	radius := bombs[x][2]
	return radius*radius >= dx*dx+dy*dy
}

// 对于一个爆炸的炸弹，可以考虑构建有向图来表示炸弹之间的引爆关系
func createGraph(bombs [][]int) map[int][]int {
	n := len(bombs)
	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j && canBomb(bombs, i, j) {
				graph[i] = append(graph[i], j)
			}
		}
	}
	return graph
}

func bfs(graph map[int][]int, start int) int {
	//记录节点是否被走过
	visited := make(map[int]bool)
	visited[start] = true
	//记录炸弹被引爆的个数
	cnt := 1
	//深度优先遍历
	queue := []int{start}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, idx := range graph[cur] {
			if !visited[idx] {
				visited[idx] = true
				cnt++
				queue = append(queue, idx)
			}
		}
	}
	return cnt
}

func maximumDetonation(bombs [][]int) int {
	n := len(bombs)
	graph := createGraph(bombs)
	//从第一个炸弹开始遍历，记录每个炸弹开始引爆能炸几个，选出其中最大的
	ans := 0
	for i := 0; i < n; i++ {
		cnt := bfs(graph, i)
		ans = max(ans, cnt)
	}
	return ans
}
