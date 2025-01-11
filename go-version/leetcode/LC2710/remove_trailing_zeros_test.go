package LC2710

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/29 下午8:41
 * @FilePath: remove_trailing_zeros_test
 * @Description:
 */

type TestData map[string]struct {
	Num       string `yaml:"Num"`
	ExpectRes string `yaml:"ExpectRes"`
}

func TestRemoveTrailingZeros(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := removeTrailingZeros(caseData.Num)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
