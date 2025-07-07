package LC1353

import (
	"container/heap"
	"sort"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/7/7 上午9:03
 * @FilePath: max_events
 * @Description: 使用最小堆
 */

func maxEvents(events [][]int) int {
	var maxDay, ans int
	for _, event := range events {
		maxDay = max(maxDay, event[1])
	}

	pq := &minHeap{}
	heap.Init(pq)
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	for i, j := 1, 0; i <= maxDay; i++ {
		for j < len(events) && events[j][0] <= i {
			heap.Push(pq, events[j][1])
			j++
		}
		for pq.Len() > 0 && (*pq)[0] < i {
			heap.Pop(pq)
		}
		if pq.Len() > 0 {
			heap.Pop(pq)
			ans++
		}
	}
	return ans
}

type minHeap []int

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
