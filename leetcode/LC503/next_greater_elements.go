package LC503

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/24 下午11:35
 * @FilePath: next_greater_elements
 * @Description: 单调栈
 */

func nextGreaterElements(nums []int) []int {
	//ans存放最终的答案
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = -1
	}
	//初始化单调栈，用于存放元素的下标
	stack := make([]int, 0)
	//循环数组构造2倍长度即可全部遍历
	for i, n := 0, len(nums); i < 2*n-1; i++ {
		//如果栈非空，且当前栈顶比当前元素还要小
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			//当前栈顶出栈，将比此栈顶更大的元素加入答案中
			ans[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		//否则的话就入栈
		stack = append(stack, i%n)
	}
	return ans
}
