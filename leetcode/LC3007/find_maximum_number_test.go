package LC3007

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/8/21 下午11:30
 * @FilePath: find_maximum_number_test
 * @Description:
 */

type TestData map[string]struct {
	K         int64 `yaml:"K"`
	X         int   `yaml:"X"`
	ExpectRes int64 `yaml:"ExpectRes"`
}

func TestFindMaximumNumber(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := findMaximumNumber(caseData.K, caseData.X)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
