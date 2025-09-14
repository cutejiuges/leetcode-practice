package LC1733

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/9/10 下午11:27
 * @FilePath: minimum_teachings
 * @Description: 找到不能沟通的好友中，会的最多的语言，把剩下的人教会就行了
 */

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	notCommunicate := make(map[int]struct{}) // 存储不能互相沟通的好友
	// 遍历好友对，找出不能沟通的好友
	for _, friendship := range friendships {
		set := make(map[int]struct{})                    // 存储当前好友对0会的语言
		var flag bool                                    // 标记当前好友对能否沟通
		for _, lan := range languages[friendship[0]-1] { // 存储0会的语言
			set[lan] = struct{}{}
		}
		for _, lan := range languages[friendship[1]-1] { // 判断1会不会0的语言
			if _, ok := set[lan]; ok {
				flag = true
				break
			}
		}
		if !flag { // 如果两个人没有共同语言，认为不能沟通
			notCommunicate[friendship[0]-1] = struct{}{}
			notCommunicate[friendship[1]-1] = struct{}{}
		}
	}
	// 遍历不能沟通的人列表，找到最多人都会的语言的数量
	var maxCnt int
	cnt := make([]int, n+1)
	for friend, _ := range notCommunicate {
		for _, lan := range languages[friend] {
			cnt[lan]++
			maxCnt = max(maxCnt, cnt[lan])
		}
	}
	// 最终教会剩下的还不会这个语言的人
	return len(notCommunicate) - maxCnt
}
