package lc2368

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

type TestData map[string]struct {
	N          int     `yaml:"N"`
	Edges      [][]int `yaml:"Edges"`
	Restricted []int   `yaml:"Restricted"`
	ExpectRes  int     `yaml:"ExpectRes"`
}

func TestReachableNodes(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for key, val := range testData {
		caseName, caseData := key, val
		t.Run(caseName, func(t *testing.T) {
			res := reachableNodes(caseData.N, caseData.Edges, caseData.Restricted)
			if res != caseData.ExpectRes {
				t.Errorf("%s, res = %d, ExpectRes = %d", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
