package lc2834

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

type TestData map[string]struct {
	N         int `yaml:"N"`
	Traget    int `yaml:"Target"`
	ExpectRes int `yaml:"ExpectRes"`
}

func TestMinimumPossibleSum(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for key, val := range testData {
		caseName, caseData := key, val
		t.Run(caseName, func(t *testing.T) {
			res := minimumPossibleSum(caseData.N, caseData.Traget)
			if res != caseData.ExpectRes {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
