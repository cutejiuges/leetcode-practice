/*
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024-03-26 08:30:16
 * @LastEditors: cutejiuge cutejiuge@163.com
 * @LastEditTime: 2024-03-27 22:33:42
 * @FilePath: /leetcode-practice/leetcode/LC2642/shortest_path.go
 * @Description: 最短路径问题，dijkstra算法
 */
package lc2642

import (
	"container/heap"
	"math"
)

// Pair 定义结构存储节点和消费，存临接表的边和结果集中的顶点
type Pair struct {
	node int
	cost int
}

type Graph struct {
	graph [][]Pair
}

// Constructor 根据边的信息构造临接表
func Constructor(n int, edges [][]int) Graph {
	graph := make([][]Pair, n)
	for i := range graph {
		graph[i] = []Pair{}
	}
	for _, e := range edges {
		start, end, cost := e[0], e[1], e[2]
		graph[start] = append(graph[start], Pair{end, cost})
	}
	return Graph{graph: graph}
}

// AddEdge 添加一条边进临接表，和构造时操作完全一致
func (this *Graph) AddEdge(edge []int) {
	start, end, cost := edge[0], edge[1], edge[2]
	(*this).graph[start] = append((*this).graph[start], Pair{end, cost})
}

// ShortestPath 求给定2个节点的最短路径
func (this *Graph) ShortestPath(node1, node2 int) int {
	//初始化优先队列存放结果集的顶点，标记所在节点和到此节点的最小消费
	pq := &PriorityQueue{}
	//标记从node1到每一个节点的最短距离
	dist := make([]int, len((*this).graph))
	for i := range dist {
		dist[i] = math.MaxInt
	}
	//到自己的最短距离肯定是0
	dist[node1] = 0
	//吧自己添加到结果集中
	heap.Push(pq, Pair{node1, 0})

	//dijkstra算法
	for len(*pq) > 0 {
		cur := heap.Pop(pq).(Pair)
		curNode, curCost := cur.node, cur.cost
		if curNode == node2 {
			return curCost
		}
		for _, e := range (*this).graph[curNode] {
			nextNode, cost := e.node, e.cost
			if dist[nextNode] > curCost+cost {
				dist[nextNode] = curCost + cost
				heap.Push(pq, Pair{nextNode, dist[nextNode]})
			}
		}
	}
	return -1
}

// PriorityQueue 实现优先队列，用于处理dijkstra算法中的最短边获取
type PriorityQueue []Pair

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq *PriorityQueue) Push(v any) {
	*pq = append(*pq, v.(Pair))
}

func (pq *PriorityQueue) Pop() any {
	n := len(*pq)
	v := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return v
}
