package LC1846

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

type TestData map[string]struct {
	Arr       []int `yaml:"arr"`
	ExpectRes int   `yaml:"expect_res"`
}

func Test_maximumElementAfterDecrementingAndRearranging(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := maximumElementAfterDecrementingAndRearranging(caseData.Arr)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
