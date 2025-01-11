package LC3304

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/9/30 下午9:21
 * @FilePath: kth_character
 * @Description:
 */

func kthCharacter(k int) byte {
	ss := []byte{'a'}
	for k > len(ss) {
		cur := ss
		for _, ch := range cur {
			c := byte('a' + (ch-'a'+1)%26)
			ss = append(ss, c)
		}
	}
	return ss[k-1]
}
