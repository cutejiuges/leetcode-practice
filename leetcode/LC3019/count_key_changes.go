package LC3019

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/1/7 下午11:25
 * @FilePath: count_key_changes
 * @Description: 按键变更的次数
 */

func countKeyChanges(s string) int {
	bytes := []byte(s)
	ans := 0
	for i := 1; i < len(bytes); i++ {
		diff := int(bytes[i-1]) - int(bytes[i])
		if diff == 32 || diff == -32 || diff == 0 {
			continue
		}
		ans++
	}
	return ans
}
