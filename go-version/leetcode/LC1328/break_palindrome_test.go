package LC1328

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/5 下午11:32
 * @FilePath: break_palindrome_test
 * @Description:
 */

type TestData map[string]struct {
	Palindrome string `yaml:"palindrome"`
	ExpectRes  string `yaml:"expect_res"`
}

func TestBreakPalindrome(t *testing.T) {
	var testdata TestData
	util.DataProvide(&testdata)

	for caseName, caseData := range testdata {
		t.Run(caseName, func(t *testing.T) {
			res := breakPalindrome(caseData.Palindrome)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
