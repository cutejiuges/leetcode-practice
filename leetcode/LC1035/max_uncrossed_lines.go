package LC1035

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/8/11 上午11:31
 * @FilePath: max_uncrossed_lines
 * @Description:
 */

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	row, col := len(nums1), len(nums2)
	dp := make([][]int, row+1)
	for i := 0; i <= row; i++ {
		dp[i] = make([]int, col+1)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if nums1[i] == nums2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[row][col]
}
