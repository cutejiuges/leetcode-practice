package main

import (
	"fmt"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/24 上午10:56
 * @FilePath: is_possible_to_rearrange
 * @Description:
 */

func isPossibleToRearrange(s, t string, k int) bool {
	subStringCnt := make(map[string]int)
	lengthOfSubstring := len(s) / k

	for i := 0; i < len(s); i += lengthOfSubstring {
		subStringCnt[s[i:i+lengthOfSubstring]]++
		subStringCnt[t[i:i+lengthOfSubstring]]--
	}

	for _, cnt := range subStringCnt {
		if cnt != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPossibleToRearrange("abcd", "cdab", 2))
	fmt.Println(isPossibleToRearrange("aabbcc", "bbaacc", 3))
	fmt.Println(isPossibleToRearrange("aabbcc", "bbaacc", 2))
	fmt.Println(isPossibleToRearrange("nc", "cn", 1))
}
