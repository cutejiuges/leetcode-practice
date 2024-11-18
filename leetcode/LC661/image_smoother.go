package LC661

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/18 下午10:55
 * @FilePath: image_smoother
 * @Description: 前缀和实现
 */

func imageSmoother(img [][]int) [][]int {
	// 初始化前缀和数组
	m, n := len(img), len(img[0])
	preSum := make([][]int, m+2)
	preSum[0] = make([]int, n+2)
	//preSum[i][j] 表示从(0,0)到(i-1, j-1)位置区块内元素之和
	for i := 1; i <= m; i++ {
		preSum[i] = make([]int, n+2)
		for j := 1; j <= n; j++ {
			preSum[i][j] = preSum[i][j-1] + preSum[i-1][j] - preSum[i-1][j-1] + img[i-1][j-1]
		}
	}
	//创建答案数组进行填充
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			left := max(0, j-1)
			right := min(n-1, j+1)
			up := max(0, i-1)
			down := min(m-1, i+1)
			sum := preSum[down+1][right+1] - preSum[down+1][left] - preSum[up][right+1] + preSum[up][left]
			ans[i][j] = sum / ((right - left + 1) * (down - up + 1))
		}
	}
	return ans
}
