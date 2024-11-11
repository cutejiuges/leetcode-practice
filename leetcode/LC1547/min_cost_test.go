package LC1547

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/11 下午11:15
 * @FilePath: min_cost_test
 * @Description:
 */

type TestData map[string]struct {
	N         int   `yaml:"n"`
	Cuts      []int `yaml:"cuts"`
	ExpectRes int   `yaml:"expect_res"`
}

func TestMinCost(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := minCost(caseData.N, caseData.Cuts)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
