package LC729

import "github.com/emirpasic/gods/trees/redblacktree"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/1/2 下午10:51
 * @FilePath: my_calendar
 * @Description: 日历会议订阅
 */

type MyCalendar struct {
	*redblacktree.Tree
}

func Constructor() MyCalendar {
	tree := redblacktree.NewWithIntComparator()
	tree.Put(-1, -1) //加入一个哨兵
	return MyCalendar{tree}
}

func (this *MyCalendar) Book(startTime int, endTime int) bool {
	floor, _ := this.Floor(startTime)
	if floor.Value.(int) > startTime {
		return false
	}
	if it := this.IteratorAt(floor); it.Next() && it.Key().(int) < endTime {
		return false
	}
	this.Put(startTime, endTime)
	return true
}
