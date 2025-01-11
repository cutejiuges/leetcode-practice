package LC1052

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

type TestData map[string]struct {
	Customers []int `yaml:"Customers"`
	Grumpy    []int `yaml:"Grumpy"`
	Minutes   int   `yaml:"Minutes"`
	ExpectRes int   `yaml:"ExpectRes"`
}

func TestMaxSatisfied(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for caseName, caseData := range testData {
		res := maxSatisfied(caseData.Customers, caseData.Grumpy, caseData.Minutes)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
