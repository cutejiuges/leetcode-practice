package LC2207

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/9/24 下午11:58
 * @FilePath: maximum_subsequence_count_test
 * @Description:
 */

type TestData map[string]struct {
	Text      string `yaml:"Text"`
	Pattern   string `yaml:"Pattern"`
	ExpectRes int64  `yaml:"ExpectRes"`
}

func TestMaximumSubsequenceCount(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := maximumSubsequenceCount(caseData.Text, caseData.Pattern)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
