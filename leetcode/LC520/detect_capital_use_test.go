package LC520

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/23 下午9:06
 * @FilePath: detect_capital_use_test
 * @Description:
 */

type TestData map[string]struct {
	Word      string `yaml:"Word"`
	ExpectRes bool   `yaml:"ExpectRes"`
}

func TestDetectCapitalUse(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := detectCapitalUse(caseData.Word)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
