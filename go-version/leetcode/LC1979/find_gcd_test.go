package LC1979

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

type TestData map[string]struct {
	Nums      []int `yaml:"nums"`
	ExpectRes int   `yaml:"expectRes"`
}

func TestFindGCD(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := findGCD(caseData.Nums)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
