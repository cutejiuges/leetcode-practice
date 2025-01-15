package lc3066

import "container/heap"

type minHeap []int

func (h *minHeap) Len() int {
	return len(*h)
}

func (h *minHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *minHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func minOperations(nums []int, k int) int {
	ans := 0
	pq := minHeap{}
	heap.Init(&pq)
	for i := range nums {
		heap.Push(&pq, nums[i])
	}

	for pq[0] < k {
		x, y := heap.Pop(&pq).(int), heap.Pop(&pq).(int)
		heap.Push(&pq, 2*x+y)
		ans++
	}
	return ans
}
