package lc1997

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

type TestData map[string]struct {
	NextVisit []int `yaml:"NextVisit"`
	ExpectRes int   `yaml:"ExpectRes"`
}

func TestFirstDayBeenInAllRooms(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for key, val := range testData {
		caseName, caseData := key, val
		res := firstDayBeenInAllRooms(caseData.NextVisit)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
