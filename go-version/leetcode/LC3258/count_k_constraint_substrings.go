package LC3258

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/12 下午10:52
 * @FilePath: count_k_constraint_substrings
 * @Description:
 */

// 使用滑动窗口的思路来实现，确定left的未知，枚举满足条件的right
func countKConstraintSubstrings(s string, k int) int {
	n := len(s)
	cnt := [2]int{0, 0} //记录滑动窗口内0、1字符的数量
	ans := 0

	//使用left和right标记滑动窗口的左右区间，右区间最大到n-1
	for left, right := 0, 0; right < n; right++ {
		cnt[s[right]-'0']++                      //窗口内右端点字符数计入
		for ; cnt[0] > k && cnt[1] > k; left++ { //如果窗口内的字符串不满足k约束，需要将左端点右移，缩小窗口
			cnt[s[left]-'0']-- //移动之前要把左端点的字符数排除
		}
		ans += right - left + 1 //这一段区间内，所有的left开始的字符串都是符合k约束的，都加到答案上
	}
	return ans
}
