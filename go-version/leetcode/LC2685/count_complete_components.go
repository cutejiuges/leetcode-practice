package LC2685

func countCompleteComponents(n int, edges [][]int) int {
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	visited := make([]bool, n)

	var dfs func(int)
	var v, e int
	dfs = func(u int) {
		if visited[u] {
			return
		}
		visited[u] = true
		v++
		e += len(graph[u])
		for _, next := range graph[u] {
			if !visited[next] {
				dfs(next)
			}
		}
	}

	var ans int
	for i := 0; i < n; i++ {
		if !visited[i] {
			v, e = 0, 0
			dfs(i)
			if v*(v-1) == e {
				ans++
			}
		}
	}
	return ans
}
