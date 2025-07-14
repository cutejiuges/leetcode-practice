package LC1290

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/7/14 下午11:24
 * @FilePath: get_decimal_value
 * @Description:
 */

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	node := head
	ans := 0
	for node != nil {
		ans *= 2
		ans += node.Val
		node = node.Next
	}
	return ans
}
