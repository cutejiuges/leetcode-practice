/*
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024-03-03 14:50:45
 * @LastEditors: cutejiuge cutejiuge@163.com
 * @LastEditTime: 2024-03-03 15:01:35
 * @FilePath: /leetcode-practice/leetcode/LC225/my_stack.go
 * @Description: leetcode 225练习，使用队列实现一个栈，没什么技巧，纯粹就是数据结构的考察
 */
package lc225

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{
		queue: []int{},
	}
}

func (stack *MyStack) Push(x int) {
	length := len(stack.queue)
	stack.queue = append(stack.queue, x)
	for i := length; i > 0; i-- {
		stack.queue = append(stack.queue, stack.queue[0])
		stack.queue = stack.queue[1:]
	}
}

func (stack *MyStack) Pop() int {
	x := stack.queue[0]
	stack.queue = stack.queue[1:]
	return x
}

func (stack *MyStack) Top() int {
	return stack.queue[0]
}

func (stack *MyStack) Empty() bool {
	return len(stack.queue) == 0
}
