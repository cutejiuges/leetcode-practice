package lc2810

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

type TestData map[string]struct {
	S         string `yaml:"S"`
	ExpectRes string `yaml:"ExpectRes"`
}

func TestFinalString(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := finalString(caseData.S)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}

		res = finalString2(caseData.S)
		if res != caseData.ExpectRes {
			t.Errorf("%s, %s, res = %v, ExpectRes = %v", caseName, "finalString2", res, caseData.ExpectRes)
		}
	}
}
