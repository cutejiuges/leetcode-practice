package LC2492

import "math"

func minScore(n int, roads [][]int) int {

	// 是否访问过该节点
	visited := make([]bool, n+1)
	graph := make([][]edge, n+1)

	// 初始化联通路径
	for _, road := range roads {
		u, v, dist := road[0], road[1], road[2]
		graph[u] = append(graph[u], edge{v, dist})
		graph[v] = append(graph[v], edge{u, dist})
	}

	ans := math.MaxInt

	var dfs func(int)
	dfs = func(u int) {
		if !visited[u] {
			visited[u] = true
		}

		for _, e := range graph[u] {
			if e.dist < ans {
				ans = e.dist
			}
			if !visited[e.v] {
				dfs(e.v)
			}
		}
	}
	dfs(1)
	return ans
}

// 定义节点之间的边类型，v指的是去向的位置，dist指的是到v的距离
type edge struct {
	v    int // 目的地
	dist int // 距离
}
