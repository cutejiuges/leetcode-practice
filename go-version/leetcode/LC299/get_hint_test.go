package lc299

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

type TestData map[string]struct {
	Secret    string `yaml:"Secret"`
	Guess     string `yaml:"Guess"`
	ExpectRes string `yaml:"ExpectRes"`
}

func TestGetHint(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for key, val := range testData {
		caseName, caseData := key, val
		t.Run(caseName, func(t *testing.T) {
			res := getHint(caseData.Secret, caseData.Guess)
			if res != caseData.ExpectRes {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
