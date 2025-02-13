package lc1742

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

type TestData map[string]struct {
	LowLimit  int `yaml:"low_limit"`
	HighLimit int `yaml:"high_limit"`
	ExpectRes int `yaml:"expect_res"`
}

func TestCountBalls(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := countBalls(caseData.LowLimit, caseData.HighLimit)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
