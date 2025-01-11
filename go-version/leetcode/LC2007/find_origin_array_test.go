package lc2007

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

type TestData map[string]struct {
	Changed   []int `yaml:"Changed"`
	ExpectRes []int `yaml:"ExpectRes"`
}

func TestFindOriginArray(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for caseName, caseData := range testData {
		res := findOriginalArray(caseData.Changed)
		if len(res) != len(caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
		for i := range res {
			if res[i] != caseData.ExpectRes[i] {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		}
	}
}
