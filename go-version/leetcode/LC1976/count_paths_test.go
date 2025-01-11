package lc1976

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

type TestData map[string]struct {
	N         int     `yaml:"N"`
	Roads     [][]int `yaml:"Roads"`
	ExpectRes int     `yaml:"ExpectRes"`
}

func TestCountPaths(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for key, val := range testData {
		caseName, caseData := key, val
		t.Run(caseName, func(t *testing.T) {
			res := countPaths(caseData.N, caseData.Roads)
			if res != caseData.ExpectRes {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
