package lc3270

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

type TestData map[string]struct {
	Num1      int `yaml:"num1"`
	Num2      int `yaml:"num2"`
	Num3      int `yaml:"num3"`
	ExpectRes int `yaml:"expect_res"`
}

func TestGenerateKey(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := generateKey(caseData.Num1, caseData.Num2, caseData.Num3)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
