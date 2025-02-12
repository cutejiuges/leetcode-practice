package lc1760

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

type TestData map[string]struct {
	Nums          []int `yaml:"nums"`
	MaxOperations int   `yaml:"max_operations"`
	ExpectRes     int   `yaml:"expect_res"`
}

func TestMinimumSize(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := minimumSize(caseData.Nums, caseData.MaxOperations)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
