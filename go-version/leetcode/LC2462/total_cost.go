package LC2462

import "container/heap"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/1 上午1:13
 * @FilePath: total_cost
 * @Description: 从数组的两头取最小的代价，考虑使用最小堆来实现
 */

type pair struct {
	val, pos int
}

type priorityQueue []pair

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	if pq[i].val == pq[j].val {
		return pq[i].pos < pq[j].pos
	}
	return pq[i].val < pq[j].val
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(v any) {
	*pq = append(*pq, v.(pair))
}
func (pq *priorityQueue) Pop() any {
	old, n := *pq, len(*pq)
	*pq = old[:n-1]
	v := old[n-1]
	return v
}

func totalCost(costs []int, k, candidates int) int64 {
	//初始化优先队列，组成最小堆
	n := len(costs)
	pq := &priorityQueue{}
	heap.Init(pq)

	//确定左边和右边即将加入队列的元素下标范围
	left, right := candidates-1, n-candidates
	//将元素加入到最小堆
	if left+1 >= right {
		//如果元素范围有相交或者即将相交，全部元素都加入最小堆
		for i := 0; i < n; i++ {
			heap.Push(pq, pair{val: costs[i], pos: i})
		}
	} else {
		//如果元素无相交，则分别将左侧和右侧元素加入最小堆
		for i := 0; i <= left; i++ {
			heap.Push(pq, pair{val: costs[i], pos: i})
		}
		for i := right; i < n; i++ {
			heap.Push(pq, pair{val: costs[i], pos: i})
		}
	}

	ans := int64(0)
	//开始进行k轮选择
	for i := k; i > 0; i-- {
		//选出当前最小的组合
		minCostPair := heap.Pop(pq).(pair)
		cost, pos := minCostPair.val, minCostPair.pos
		//将代价加入到答案中
		ans += int64(cost)
		//如果元素没有相交，需要加入新的元素
		if left+1 < right {
			if pos <= left {
				//如果本次选中左侧元素，那么还是需要左侧元素入队
				left++
				heap.Push(pq, pair{val: costs[left], pos: left})
			} else {
				//如果本次选中右侧元素，那么还是需要右侧元素入队
				right--
				heap.Push(pq, pair{val: costs[right], pos: right})
			}
		}
	}
	return ans
}
