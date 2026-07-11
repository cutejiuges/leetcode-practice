package LC2685

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

type TestData map[string]struct {
	N         int     `yaml:"n"`
	Edges     [][]int `yaml:"edges"`
	ExpectRes int     `yaml:"expectRes"`
}

func TestCountCompleteComponents(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := countCompleteComponents(caseData.N, caseData.Edges)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
