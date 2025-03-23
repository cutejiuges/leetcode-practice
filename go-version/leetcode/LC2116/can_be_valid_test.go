package LC2116

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/23 上午11:38
 * @FilePath: can_be_valid_test
 * @Description:
 */

type TestData map[string]struct {
	S         string `yaml:"s"`
	Locked    string `yaml:"locked"`
	ExpectRes bool   `yaml:"expect_res"`
}

func TestCanBeValid(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := canBeValid(caseData.S, caseData.Locked)
			if res != caseData.ExpectRes {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
