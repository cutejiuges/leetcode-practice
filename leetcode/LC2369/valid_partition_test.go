package lc2369

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

type TestData map[string]struct {
	Nums      []int `yaml:"Nums"`
	ExpectRes bool  `yaml:"ExpectRes"`
}

func TestValidPartition(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for key, val := range testData {
		caseName, caseData := key, val
		t.Run(caseName, func(t *testing.T) {
			res := validPartition(caseData.Nums)
			if res != caseData.ExpectRes {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
