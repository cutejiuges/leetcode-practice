package LC3532

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

type TestData map[string]struct {
	N         int     `yaml:"n"`
	Nums      []int   `yaml:"nums"`
	MaxDiff   int     `yaml:"maxDiff"`
	Queries   [][]int `yaml:"queries"`
	ExpectRes []bool  `yaml:"expectRes"`
}

func TestPathExistenceQueries(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := pathExistenceQueries(caseData.N, caseData.Nums, caseData.MaxDiff, caseData.Queries)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
