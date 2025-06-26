package LC2311

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/6/26 下午11:19
 * @FilePath: longest_subsequence_test
 * @Description:
 */

type TestData map[string]struct {
	S         string `yaml:"s"`
	K         int    `yaml:"k"`
	ExpectRes int    `yaml:"expect_res"`
}

func TestLongestSubsequence(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := longestSubsequence(caseData.S, caseData.K); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
