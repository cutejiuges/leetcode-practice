package lc2312

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

type TestData map[string]struct {
	M         int     `yaml:"M"`
	N         int     `yaml:"N"`
	Prices    [][]int `yaml:"Prices"`
	ExpectRes int     `yaml:"ExpectRes"`
}

func TestSellingWood(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for key, val := range testData {
		caseName, caseData := key, val
		t.Run(caseName, func(t *testing.T) {
			res := sellingWood(caseData.M, caseData.N, caseData.Prices)
			if res != int64(caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
