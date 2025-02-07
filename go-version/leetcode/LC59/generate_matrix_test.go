package lc59

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

type TestData map[string]struct {
	N         int     `yaml:"n"`
	ExpectRes [][]int `yaml:"expect_res"`
}

func TestGenerateMatrix(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := generateMatrix(caseData.N)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
