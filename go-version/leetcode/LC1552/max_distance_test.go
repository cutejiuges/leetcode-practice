package lc1552

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

type TestData map[string]struct {
	Position  []int `yaml:"position"`
	M         int   `yaml:"m"`
	ExpectRes int   `yaml:"expect_res"`
}

func TestMaxDistance(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := maxDistance(caseData.Position, caseData.M)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData)
		}
	}
}
