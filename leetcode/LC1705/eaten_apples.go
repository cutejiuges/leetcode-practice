package LC1705

import (
	"container/heap"
	datastruct "cutejiuge/leetcode-practice/data_struct"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/24 下午10:39
 * @FilePath: eaten_apples
 * @Description: 用最小堆来维护腐烂时间
 */

type pair struct {
	num, rotDay int
}

func eatenApples(apples, days []int) int {
	ans := 0
	minHeap := datastruct.PriorityQueue[pair]{}

	for i := 0; i < len(apples) || minHeap.Len() > 0; i++ {
		for minHeap.Len() > 0 && minHeap[0].Value.rotDay == i {
			//队首的苹果已经腐烂了
			heap.Pop(&minHeap)
		}
		if i < len(apples) && apples[i] > 0 {
			//有效的苹果数，加入队列
			heap.Push(&minHeap, &datastruct.Item[pair]{Value: pair{apples[i], i + days[i]}, Priority: i + days[i]})
		}
		if minHeap.Len() > 0 {
			//吃掉一个苹果
			ans++
			minHeap[0].Value.num--
			if minHeap[0].Value.num == 0 {
				heap.Pop(&minHeap)
			}
		}
	}
	return ans
}
