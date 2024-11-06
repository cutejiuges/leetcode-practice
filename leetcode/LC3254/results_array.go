package LC3254

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/6 下午11:11
 * @FilePath: results_array
 * @Description: 遍历计算即可
 */

func resultsArray(nums []int, k int) []int {
	n := len(nums)
	res := make([]int, n-k+1)

	//使用-1初始化答案数组
	for i := range res {
		res[i] = -1
	}

	//枚举每一个长度为k的子数组
	for i := 0; i <= n-k; i++ {
		flag := true //标记当前子数组数否满足要求
		//从起始位置枚举k个连续元素
		for j := i + 1; j < i+k; j++ {
			if nums[j]-nums[j-1] != 1 { //如果不满足元素连续的要求
				flag = false //标记次子数组不满足要求
				break        //跳出枚举
			}
		}
		if flag {
			res[i] = nums[i+k-1]
		} //如果是满足条件的子数组，取其中的最大值
	}
	return res
}
