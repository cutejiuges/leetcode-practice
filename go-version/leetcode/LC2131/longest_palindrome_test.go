package LC2131

import (
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/25 下午7:00
 * @FilePath: longest_palindrome_test.go
 * @Description:
 */

type TestData map[string]struct {
	Words     []string `yaml:"words"`
	ExpectRes int      `yaml:"expect_res"`
}

func TestLongestPalindrome(t *testing.T) {
	var testData TestData
	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := longestPalindrome(caseData.Words); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
