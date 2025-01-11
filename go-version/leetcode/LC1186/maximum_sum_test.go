package LC1186

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/7/21 下午3:01
 * @FilePath: maximum_sum_test
 * @Description:
 */

type TestData map[string]struct {
	Arr       []int `yaml:"Arr"`
	ExpectRes int   `yaml:"ExpectRes"`
}

func TestMaximumSum(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := maximumSum2(caseData.Arr)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectREs = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
