/*
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024-03-05 23:10:31
 * @LastEditors: cutejiuge cutejiuge@163.com
 * @LastEditTime: 2024-03-06 00:57:59
 * @FilePath: /leetcode-practice/leetcode/LC1976/count_paths.go
 * @Description: leetcode 1976练习，考虑使用dijkstra算法来实现最少花费的路径求解
 * 同时，为了记录做短路径的数量，还需要考虑将路径数量单独存储起来
 */
package lc1976

import (
	"container/heap"
	"math"
)

// 存储边结构，存放下一个节点的编号和当前节点到达下一个节点的距离
type edge struct {
	dest, cost int
}

// 记录一个节点信息，存放从0节点到dijkstra结果集中节点的距离和此节点编号
type pair struct {
	distance int64
	cur      int
}

// 优先队列，存储结果集，即到达节点n途经的节点
type priorityQueue []pair

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq *priorityQueue) Push(val any) {
	*pq = append(*pq, val.(pair))
}

func (pq *priorityQueue) Pop() any {
	n := len(*pq)
	val := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return val
}

func countPaths(n int, roads [][]int) int {
	const mod = int64(1e9 + 7)
	//转换成临接表
	graph := make([][]edge, n)
	for _, road := range roads {
		graph[road[0]] = append(graph[road[0]], edge{road[1], road[2]})
		graph[road[1]] = append(graph[road[1]], edge{road[0], road[2]})
	}

	//新建一个列表，存储从0开始到节点i的最短距离
	dist := make([]int64, n)
	for i := range dist {
		dist[i] = math.MaxInt64
	}
	//新建一个列表，存储从0开始到节点i的最短路径条数
	cnt := make([]int64, n)

	//初始化存储结构和数据
	queue := &priorityQueue{{0, 0}} //初始节点从0开始
	heap.Init(queue)                //初始化堆
	cnt[0] = 1                      //0到0的有一条路
	dist[0] = 0                     //0到0的距离一定是0

	//dijkstra算法，取出结果集中到当前节点最小的节点记录
	for len(*queue) > 0 {
		//弹出队列中dist距离最小的节点
		minNode := heap.Pop(queue).(pair)
		//如果结果集中记录的最小距离已经比实际距离更大，不进行本次循环
		distance, cur := minNode.distance, minNode.cur
		if distance > dist[cur] {
			continue
		}

		//从这个最小节点，开始找最近的下一个节点
		for _, next := range graph[cur] {
			v, cost := next.dest, next.cost
			if distance+int64(cost) < dist[v] {
				//如果到下一节点的距离+当前节点最短距离小于目前记录的到下一节点的最小距离，则更新最小距离，并将下一节点信息写入结果集中
				dist[v] = distance + int64(cost)
				heap.Push(queue, pair{distance + int64(cost), v})
				cnt[v] = cnt[cur]
			} else if distance+int64(cost) == dist[v] {
				//如果相等的话则更新一下到下一节点的最短路径数量
				cnt[v] = (cnt[cur] + cnt[v]) % mod
			}
		}
	}
	return int(cnt[n-1])
}
