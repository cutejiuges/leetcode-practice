package lc2129

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

type TestData map[string]struct {
	Title     string `yaml:"Title"`
	ExpectRes string `yaml:"ExpectRes"`
}

func TestCapitalizeTitle(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for key, val := range testData {
		caseName, caseData := key, val
		t.Run(caseName, func(t *testing.T) {
			res := capitalizeTitle(caseData.Title)
			if res != caseData.ExpectRes {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
