package lc331

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

type TestData map[string]struct {
	Preorder  string `yaml:"Preorder"`
	ExpectRes bool   `yaml:"ExpectRes"`
}

func TestIsValidSerialization(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for caseName, caseData := range testData {
		res := isValidSerialization(caseData.Preorder)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
