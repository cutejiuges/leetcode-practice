package LC1766

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

type TestData map[string]struct {
	Nums      []int   `yaml:"Nums"`
	Edges     [][]int `yaml:"Edges"`
	ExpectRes []int   `yaml:"ExpectRes"`
}

func TestGetCoprimes(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for caseName, caseData := range testData {
		res := getCoprimes(caseData.Nums, caseData.Edges)
		var flag = true
		if len(res) != len(caseData.ExpectRes) {
			flag = false
		}
		for i := range res {
			if res[i] != caseData.ExpectRes[i] {
				flag = false
			}
		}
		if !flag {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
