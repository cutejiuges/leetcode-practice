package LC1652

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/5 上午12:22
 * @FilePath: decrypt
 * @Description: 循环+滑动窗口
 */

func decrypt(code []int, k int) []int {
	res := make([]int, len(code))
	if k == 0 {
		//k=0时，全部返回0即可
		return res
	}
	//逐个计算相加结果
	for i := range code {
		sum := 0
		//取相邻的k个元素
		for j := 0; j < abs(k); j++ {
			idx := 0
			//如果k>0向后取, k<0向前取
			if k > 0 {
				//循环数组，需要取模
				idx = (i + j + 1) % len(code)
			} else {
				//避免负数，需要加上一个长度
				idx = (i - j - 1 + len(code)) % len(code)
			}
			//元素求和
			sum += code[idx]
		}
		//将结果放入数组
		res[i] = sum
	}
	return res
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
