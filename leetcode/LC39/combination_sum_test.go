package LC39

import (
	"cutejiuge/leetcode-practice/util"
	"sort"
	"testing"
)

type TestData map[string]struct {
	Candidates []int   `yaml:"Candidates"`
	Target     int     `yaml:"Target"`
	ExpectRes  [][]int `yaml:"ExpectRes"`
}

func TestCombinationSum(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for caseName, caseData := range testData {
		res := combinationSum(caseData.Candidates, caseData.Target)
		if len(res) > 0 {
			sort.Slice(res, func(i, j int) bool {
				if len(res[i]) == len(res[j]) {
					return res[i][0] < res[j][0]
				}
				return len(res[i]) < len(res[j])
			})
		}
		if len(caseData.ExpectRes) > 0 {
			sort.Slice(caseData.ExpectRes, func(i, j int) bool {
				if len(caseData.ExpectRes[i]) == len(caseData.ExpectRes[j]) {
					return caseData.ExpectRes[i][0] < caseData.ExpectRes[j][0]
				}
				return len(caseData.ExpectRes[i]) < len(caseData.ExpectRes[j])
			})
		}
		if len(res) != len(caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
		for i := range res {
			if len(res[i]) != len(caseData.ExpectRes[i]) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
			for j := range res[i] {
				if res[i][j] != caseData.ExpectRes[i][j] {
					t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
				}
			}
		}
	}
}
