package LC3340

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/14 下午11:18
 * @FilePath: is_balanced
 * @Description: 奇偶数遍历放在一个整数里面，判断最后是否是0即可
 */

func isBalanced(num string) bool {
	cnt := 0
	for i, ch := range num {
		if i&1 == 0 {
			cnt += int(ch - '0')
		} else {
			cnt -= int(ch - '0')
		}
	}
	return cnt == 0
}
