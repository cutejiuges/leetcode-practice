package LC3904

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

type TestData map[string]struct {
	Nums      []int `yaml:"nums"`
	K         int   `yaml:"k"`
	ExpectRes int   `yaml:"expectRes"`
}

func TestFirstStableIndex(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := firstStableIndex(caseData.Nums, caseData.K)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
