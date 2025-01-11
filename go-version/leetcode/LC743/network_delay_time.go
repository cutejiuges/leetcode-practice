package LC743

import "container/heap"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/25 下午10:19
 * @FilePath: network_delay_time
 * @Description:
 */

type pair struct {
	dist, next int
}

type priorityQueue []pair

func (pq *priorityQueue) Len() int           { return len(*pq) }
func (pq *priorityQueue) Less(i, j int) bool { return (*pq)[i].dist < (*pq)[j].dist }
func (pq *priorityQueue) Swap(i, j int)      { (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i] }
func (pq *priorityQueue) Push(x any)         { *pq = append(*pq, x.(pair)) }
func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}

func networkDelayTime(times [][]int, n int, k int) int {
	type edge struct {
		to, cost int //定义边结构，标记终点和代价，用于构建邻接表
	}
	//构建初始化邻接表，为了保证位序，节点标记都减掉1,和下标对其
	graph := make([][]edge, n)
	for _, time := range times {
		graph[time[0]-1] = append(graph[time[0]-1], edge{time[1] - 1, time[2]})
	}

	//初始化到每个节点之间的距离是无穷大
	inf := 0x3f3f3f3f
	dist := make([]int, n)
	for i := range dist {
		dist[i] = inf
	}
	dist[k-1] = 0 //第k个节点到自己的距离是0,直接可达
	queue := &priorityQueue{{0, k - 1}}
	heap.Init(queue)
	for queue.Len() > 0 {
		cur := heap.Pop(queue).(pair)
		next := cur.next
		if dist[next] < cur.dist {
			continue
		}
		for _, e := range graph[next] {
			to := e.to
			if d := dist[next] + e.cost; d < dist[to] {
				dist[to] = d
				heap.Push(queue, pair{d, to})
			}
		}
	}

	ans := 0
	for _, d := range dist {
		if d == inf {
			return -1
		}
		ans = max(ans, d)
	}
	return ans
}
