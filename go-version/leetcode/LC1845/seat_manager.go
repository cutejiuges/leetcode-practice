package LC1845

import "container/heap"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/9/30 上午12:08
 * @FilePath: seat_manager
 * @Description: 最小堆的实现
 */

type SeatManager struct {
	pq *priorityQueue
}

func Constructor(n int) SeatManager {
	h := &priorityQueue{}
	heap.Init(h)
	for i := 1; i <= n; i++ {
		heap.Push(h, i)
	}
	return SeatManager{pq: h}
}

func (this *SeatManager) Reserve() int {
	return heap.Pop(this.pq).(int)
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(this.pq, seatNumber)
}

// 小根堆的实现
type priorityQueue []int

func (pq *priorityQueue) Len() int {
	return len(*pq)
}

func (pq *priorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *priorityQueue) Less(i, j int) bool {
	return (*pq)[i] < (*pq)[j]
}

func (pq *priorityQueue) Push(x any) {
	*pq = append(*pq, x.(int))
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[:n-1]
	return x
}
