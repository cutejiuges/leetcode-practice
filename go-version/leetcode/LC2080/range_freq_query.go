package lc2080

type RangeFreqQuery struct {
	occurrence map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	mp := make(map[int][]int)
	for i, num := range arr {
		mp[num] = append(mp[num], i)
	}
	return RangeFreqQuery{occurrence: mp}
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	position := this.occurrence[value]
	low, high := this.lowBound(position, left), this.highBound(position, right)
	return high - low
}

func (this *RangeFreqQuery) lowBound(position []int, low int) int {
	left, right := 0, len(position)-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if position[mid] < low {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

func (this *RangeFreqQuery) highBound(position []int, high int) int {
	left, right := 0, len(position)-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if position[mid] <= high {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
