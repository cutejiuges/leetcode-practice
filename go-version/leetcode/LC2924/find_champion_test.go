package LC2924

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

type TestData map[string]struct {
	N         int     `yaml:"N"`
	Edges     [][]int `yaml:"Edges"`
	ExpectRes int     `yaml:"ExpectRes"`
}

func TestFindChampion(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for caseName, caseData := range testData {
		res := findChampion(caseData.N, caseData.Edges)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}

func TestFindChampion2(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for caseName, caseData := range testData {
		res := findChampion2(caseData.N, caseData.Edges)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
