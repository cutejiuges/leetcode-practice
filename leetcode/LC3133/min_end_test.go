package LC3133

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/8/22 下午11:17
 * @FilePath: min_end_test
 * @Description:
 */

type TestData map[string]struct {
	N         int   `yaml:"N"`
	X         int   `yaml:"X"`
	ExpectRes int64 `yaml:"ExpectRes"`
}

func TestMinEnd(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := minEnd(caseData.N, caseData.X)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectREs = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
