package LC2923

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

type TestData map[string]struct {
	Grid      [][]int `yaml:"Grid"`
	ExpectRes int     `yaml:"ExpectRes"`
}

func TestFindChampion(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for caseName, caseData := range testData {
		res := findChampion(caseData.Grid)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}

func TestFindChampion2(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for caseName, caseData := range testData {
		res := findChampion2(caseData.Grid)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
