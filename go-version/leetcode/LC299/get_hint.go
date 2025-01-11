/*
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024-03-10 08:20:27
 * @LastEditors: cutejiuge cutejiuge@163.com
 * @LastEditTime: 2024-03-10 08:29:19
 * @FilePath: /leetcode-practice/leetcode/LC299/get_hint.go
 * @Description: lc299题，统计数字遍历即可
 */
package lc299

import "fmt"

func getHint(secret, guess string) string {
	bulls, cows := 0, 0
	cntSecret, cntGuess := make([]int, 10), make([]int, 10)
	for i := range secret {
		if secret[i] == guess[i] {
			bulls++
		} else {
			cntGuess[guess[i]-'0']++
			cntSecret[secret[i]-'0']++
		}
	}
	for i := 0; i < 10; i++ {
		cows += min(cntGuess[i], cntSecret[i])
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
