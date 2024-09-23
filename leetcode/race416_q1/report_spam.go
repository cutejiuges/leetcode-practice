package race416_q1

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/9/22 上午10:32
 * @FilePath: report_spam
 * @Description: 使用集合进行模拟
 */

func reportSpam(message []string, bannedWords []string) bool {
	msgSet := make(map[string]struct{})
	for _, msg := range bannedWords {
		msgSet[msg] = struct{}{}
	}

	cnt := 0
	for _, word := range message {
		if _, ok := msgSet[word]; ok {
			cnt++
		}
	}
	if cnt >= 2 {
		return true
	}
	return false
}
