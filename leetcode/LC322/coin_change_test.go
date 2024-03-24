package lc322

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

type TestData map[string]struct {
	Coins     []int `yaml:"Coins"`
	Amount    int   `yaml:"Amount"`
	ExpectRes int   `yaml:"ExpectRes"`
}

func TestCoinChange(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for key, val := range testData {
		caseName, caseData := key, val
		res := coinChange(caseData.Coins, caseData.Amount)
		if res != caseData.ExpectRes {
			t.Errorf("%s, ExpectRes = %v, res = %v \n", caseName, caseData.ExpectRes, res)
		}
	}
}
