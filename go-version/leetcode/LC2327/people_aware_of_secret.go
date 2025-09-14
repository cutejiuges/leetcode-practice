package LC2327

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/9/9 下午11:33
 * @FilePath: people_aware_of_secret
 * @Description: 两个双端队列进行模拟
 */

func peopleAwareOfSecret(n int, delay int, forget int) int {
	mod := int(1e9 + 7)
	// 两个双端队列进行模拟，一个特表示仅知道秘密，一个表示可分享秘密
	know, share := make([][2]int, 0), make([][2]int, 0)
	// 第1天只有1个人只知道秘密
	know = append(know, [2]int{1, 1})
	// 两个整数，分别表示只知道秘密的人数和能传播秘密的人数
	knowCnt, shareCnt := 1, 0

	// 从第2天开始模拟
	for i := 2; i <= n; i++ {
		// 如果当前know最早的项已经到了可以分享的时间，转移到分享队列，并减少knowCnt
		if len(know) > 0 && know[0][0] == i-delay {
			head := know[0]
			know = know[1:]
			knowCnt = (knowCnt - head[1] + mod) % mod
			shareCnt = (shareCnt + head[1]) % mod
			share = append(share, head)
		}
		// 如果当前share最早的项到了忘记的时间，删除，并减少shareCnt
		if len(share) > 0 && share[0][0] == i-forget {
			head := share[0]
			share = share[1:]
			shareCnt = (shareCnt - head[1] + mod) % mod
		}
		// 不是以上两种情况，就进行传播
		if len(share) > 0 {
			knowCnt = (knowCnt + shareCnt) % mod
			know = append(know, [2]int{i, shareCnt})
		}
	}
	// 最终人数是两部分相加
	return (knowCnt + shareCnt) % mod
}
