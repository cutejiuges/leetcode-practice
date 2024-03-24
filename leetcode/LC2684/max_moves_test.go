package lc2684

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

type TestData map[string]struct {
	Grid      [][]int `yaml:"Grid"`
	ExpectRes int     `yaml:"ExpectRes"`
}

func TestMaxMoves(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for key, val := range testData {
		caseName, caseData := key, val
		res := maxMoves(caseData.Grid)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
