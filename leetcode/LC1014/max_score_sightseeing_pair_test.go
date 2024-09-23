package LC1014

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/9/23 下午11:17
 * @FilePath: max_score_sightseeing_pair_test
 * @Description:
 */

type TestData map[string]struct {
	Values    []int `yaml:"Values"`
	ExpectRes int   `yaml:"ExpectRes"`
}

func TestMaxScoreSightseeingPair(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := maxScoreSightseeingPair(caseData.Values)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
