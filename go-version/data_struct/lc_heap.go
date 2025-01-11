package datastruct

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/24 下午10:40
 * @FilePath: lc_heap
 * @Description: 实现一个通用的heap结构，不用每次都重写了
 */

type Item[E comparable] struct {
	Value    E
	Priority int
}

type PriorityQueue[E comparable] []*Item[E]

func (pq *PriorityQueue[E]) Len() int {
	return len(*pq)
}

func (pq *PriorityQueue[E]) Less(i, j int) bool {
	return (*pq)[i].Priority < (*pq)[j].Priority
}

func (pq *PriorityQueue[E]) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue[E]) Push(x any) {
	*pq = append(*pq, x.(*Item[E]))
}

func (pq *PriorityQueue[E]) Pop() any {
	old := *pq
	n := len(old)
	val := old[n-1]
	*pq = old[:n-1]
	return val
}
