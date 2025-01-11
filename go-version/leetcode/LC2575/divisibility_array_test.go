package lc2575

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

type TestData map[string]struct {
	Word      string `yaml:"Word"`
	M         int    `yaml:"M"`
	ExpectRes []int  `yaml:"ExpectRes"`
}

func TestDivisibilityArray(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for key, val := range testData {
		caseName, caseData := key, val
		t.Run(caseName, func(t *testing.T) {
			res := divisibilityArray(caseData.Word, caseData.M)
			if len(res) != len(caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
			for i := range res {
				if res[i] != caseData.ExpectRes[i] {
					t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
				}
			}
		})
	}
}
