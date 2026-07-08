package LC3756

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

type TestData map[string]struct {
	S         string  `yaml:"s"`
	Queries   [][]int `yaml:"queries"`
	ExpectRes []int   `yaml:"expect_res"`
}

func TestSumAndMultiply(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := sumAndMultiply(caseData.S, caseData.Queries)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
