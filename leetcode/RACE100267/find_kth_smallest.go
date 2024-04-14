package RACE100267

import (
	"container/heap"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/14 上午11:15
 * @FilePath: find_kth_smallest
 * @Description: kth最小
 */

// 优先队列
type priorityQueue []int64

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq *priorityQueue) Push(val any) {
	*pq = append(*pq, val.(int64))
}

func (pq *priorityQueue) Pop() any {
	n := len(*pq)
	val := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return val
}

func findKthSmallest(coins []int, k int) int64 {
	mp := map[int64]struct{}{}
	length := len(coins)
	queue := &priorityQueue{}
	heap.Init(queue)
	for i := 0; i < length; i++ {
		for j := 1; j <= k; j++ {
			mp[int64(coins[i]*j)] = struct{}{}
		}
	}
	for key := range mp {
		heap.Push(queue, key)
	}
	ans := int64(0)
	for k > 0 && len(*queue) > 0 {
		ans = heap.Pop(queue).(int64)
		k--
	}
	return ans
}
