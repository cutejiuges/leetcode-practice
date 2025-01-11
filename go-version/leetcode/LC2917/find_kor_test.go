package lc2917

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

type TestData map[string]struct {
	Nums      []int `yaml:"Nums"`
	K         int   `yaml:"K"`
	ExpectRes int   `yaml:"ExpectRes"`
}

func TestFindKOr(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for key, val := range testData {
		caseName, caseData := key, val
		t.Run(caseName, func(t *testing.T) {
			res := findKOr(caseData.Nums, caseData.K)
			if res != caseData.ExpectRes {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
