package q3

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/11 上午11:21
 * @FilePath: max_score
 * @Description:
 */

func maxScore(n int, edges [][]int) int64 {
	graph := buildGraph(n, edges)
	visited := make([]bool, n)
	var totalSum int64
	for i := 0; i < n; i++ {
		if !visited[i] {
			if len(graph[i]) == 1 {
				totalSum += handleChain(graph, i, visited)
			} else if len(graph[i]) == 2 {
				totalSum += handleCycle(graph, i, visited)
			}
		}
	}
	return totalSum
}

// 构建邻接表
func buildGraph(n int, edges [][]int) [][]int {
	graph := make([][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	return graph
}

func handleChain(graph [][]int, start int, visited []bool) int64 {
	chain := []int{start}
	visited[start] = true
	cur := start
	for len(graph[cur]) == 2 {
		for _, neighbor := range graph[cur] {
			if !visited[neighbor] {
				visited[neighbor] = true
				chain = append(chain, neighbor)
				cur = neighbor
				break
			}
		}
	}
	var sum int64
	for i, j := 0, len(chain)-1; i < j; i, j = i+1, j-1 {
		sum += int64(i+1) * int64(j+1)
	}
	for i := 1; i < len(chain)-1; i++ {
		sum += int64(i+1) * int64(i+1)
	}
	return sum
}

func handleCycle(graph [][]int, start int, visited []bool) int64 {
	cycle := []int{start}
	visited[start] = true
	cur := start
	pre := -1
	startIdx := 0
	for {
		found := false
		for _, neighbor := range graph[cur] {
			if !visited[neighbor] {
				visited[neighbor] = true
				cycle = append(cycle, neighbor)
				pre = cur
				cur = neighbor
				found = true
				break
			} else if neighbor != pre {
				cycle = append(cycle, neighbor)
				cur = neighbor
				found = true
				break
			}
		}
		if !found {
			break
		}
		if cur == start {
			startIdx = len(cycle) - 1
			break
		}
	}
	newCycle := make([]int, len(cycle)-startIdx)
	copy(newCycle, cycle[startIdx:])
	var sum int64
	for i, j := 0, len(cycle)-1; i < j; i, j = i+1, j-1 {
		sum += int64(i+1) * int64(j+1)
	}
	for i := 1; i < len(cycle)-1; i++ {
		sum += int64(i+1) * int64(i+1)
	}
	return sum
}
