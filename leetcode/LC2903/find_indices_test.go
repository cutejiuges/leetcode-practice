package LC2903

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/19 下午11:37
 * @FilePath: find_indices_test
 * @Description:
 */

type TestData map[string]struct {
	Nums            []int `yaml:"Nums"`
	IndexDifference int   `yaml:"IndexDifference"`
	ValueDifference int   `yaml:"ValueDifference"`
	ExpectRes       []int `yaml:"ExpectRes"`
}

func TestFindIndices(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := findIndices(caseData.Nums, caseData.IndexDifference, caseData.ValueDifference)
		if res[0] != caseData.ExpectRes[0] || res[1] != caseData.ExpectRes[1] {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
