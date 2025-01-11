package LC3154

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/8/20 下午10:55
 * @FilePath: ways_to_reach_stair_test
 * @Description:
 */

type TestData map[string]struct {
	K         int `yaml:"K"`
	ExpectRes int `yaml:"ExpectRes"`
}

func TestWaysToReachStair(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := waysToReachStair(caseData.K)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
