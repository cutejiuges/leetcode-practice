package LC857

import (
	"container/heap"
	"math"
	"sort"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/2 上午12:18
 * @FilePath: mincost_to_hire_workers
 * @Description: 雇佣工人的最低代价，仍然考虑用堆来实现
 */
type priorityQueue []int

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(v any) {
	*pq = append(*pq, v.(int))
}

func (pq *priorityQueue) Pop() any {
	old, n := *pq, len(*pq)
	*pq = old[:n-1]
	v := old[n-1]
	return v
}

func mincostToHireWorkers(quality, wage []int, k int) float64 {
	n := len(quality)
	h := make([]int, n)
	for i := range h {
		h[i] = i
	}
	sort.Slice(h, func(i, j int) bool {
		a, b := h[i], h[j]
		return quality[a]*wage[b] > quality[b]*wage[a]
	})
	totalQuality := 0
	pq := &priorityQueue{}
	for i := 0; i < k-1; i++ {
		totalQuality += quality[h[i]]
		heap.Push(pq, quality[h[i]])
	}
	ans := 1e9
	for i := k - 1; i < n; i++ {
		idx := h[i]
		totalQuality += quality[idx]
		heap.Push(pq, quality[idx])
		ans = math.Min(ans, float64(wage[idx])/float64(quality[idx])*float64(totalQuality))
		totalQuality -= heap.Pop(pq).(int)
	}
	return ans
}
