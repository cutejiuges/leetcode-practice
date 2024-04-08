package LC2009

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

type TestData map[string]struct {
	Nums      []int `yaml:"Nums"`
	ExpectRes int   `yaml:"ExpectRes"`
}

func TestMinOperations(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for caseName, caseData := range testData {
		res := minOperations(caseData.Nums)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
