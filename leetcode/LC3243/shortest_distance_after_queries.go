package LC3243

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/19 下午10:57
 * @FilePath: shortest_distance_after_queries
 * @Description: 通过bfs实现
 */

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	//通过查询关系构建邻接表
	graph := make([][]int, n)
	for i := 0; i < n-1; i++ {
		graph[i] = append(graph[i], i+1)
	}
	var ans []int
	//加路径之后直接计算一次最短路径
	for _, query := range queries {
		graph[query[0]] = append(graph[query[0]], query[1])
		ans = append(ans, bfs(n, graph))
	}
	return ans
}

func bfs(n int, graph [][]int) int {
	inf := 0x3f3f3f3f      //无穷大长度，表示不可达
	dist := make([]int, n) //0到每一个节点的距离
	for i := range dist {
		dist[i] = inf //初始化为无穷大，标记不可达
	}
	dist[0] = 0      //0到0的距离是0
	q := []int{0}    //初始状态可以直接到到0节点
	for len(q) > 0 { //未访问过的节点
		cur := q[0]
		q = q[1:]                         //取出队首元素
		for _, next := range graph[cur] { //找出该节点的相邻节点
			if dist[next] < inf { //如果已经访问过该节点，本次不再操作
				continue
			}
			dist[next] = min(dist[next], dist[cur]+1) //更新节点距离
			q = append(q, next)                       //标记为已访问节点
		}
	}
	return dist[n-1]
}
