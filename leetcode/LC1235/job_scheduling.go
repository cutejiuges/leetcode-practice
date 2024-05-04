package LC1235

import "sort"

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/4 下午9:51
 * @FilePath: job_scheduling
 * @Description: 贪心+动态规划
 */

/*
 首先考虑按照会议问题，将兼职工作按照结束时间升序排列，确保可以尽可能多的选择兼职工作；
 排序之后，dp[i]表示当前选择之后的最大薪酬，对于i工作，可以不选择这一项工作，或者将其接在上一个结束的工作之后
 上一个结束的工作，可以在序列中二分查找，这里指的是，上一个endTime在当前startTime之前的工作
 选择出来之后，dp[i] = max(dp[i-1], dp[k]+profit[i])
*/

type jobInfo struct {
	StartTime, EndTime, Profit int
}

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	jobs := make([]jobInfo, n)
	for i := 0; i < n; i++ {
		jobs[i] = jobInfo{StartTime: startTime[i], EndTime: endTime[i], Profit: profit[i]}
	}
	//将序列按照结束时间升序排列
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].EndTime < jobs[j].EndTime
	})

	//进行动态规划的递推
	dp := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		pos := sort.Search(i, func(j int) bool {
			return jobs[j].EndTime > jobs[i-1].StartTime
		})
		dp[i] = max(dp[i-1], dp[pos]+jobs[i-1].Profit)
	}
	return dp[n]
}
