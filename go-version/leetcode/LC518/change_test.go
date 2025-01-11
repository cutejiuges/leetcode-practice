package lc518

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

type TestData map[string]struct {
	Amount    int   `yaml:"Amount"`
	Coins     []int `yaml:"Coins"`
	ExpectRes int   `yaml:"ExpectRes"`
}

func TestChange(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for key, val := range testData {
		caseName, caseData := key, val
		res := change(caseData.Amount, caseData.Coins)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
