package LC2306

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/9/25 上午12:10
 * @FilePath: distinct_names
 * @Description:
 */

// 后缀优化
func distinctNames(ideas []string) (res int64) {
	m := map[string]int{}
	for _, s := range ideas {
		//按照后缀来分组，状态压缩
		m[s[1:]] |= 1 << (s[0] - 'a')
	}
	//count[i][j] 表示首字母不含i但是含j的个数
	count := [26][26]int64{}
	for _, v := range m {
		for i := 0; i < 26; i++ {
			//第i位为0
			if (v>>i)&1 == 0 {
				for j := 0; j < 26; j++ {
					//第j位为1
					if (v>>j)&1 == 1 {
						count[i][j]++
					}
				}
				//第i位为1
			} else {
				for j := 0; j < 26; j++ {
					//第j位为0
					if (v>>j)&1 == 0 {
						//结算
						res += count[i][j]
					}
				}
			}
		}
	}
	//算上两个单词前后顺序的排列
	return 2 * res
}
