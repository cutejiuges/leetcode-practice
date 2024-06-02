package LC575

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/2 下午8:48
 * @FilePath: distribute_candies_test
 * @Description:
 */

type TestData map[string]struct {
	CandyType []int `yaml:"CandyType"`
	ExpectRes int   `yaml:"ExpectRes"`
}

func TestDistributeCandies(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := distributeCandies(caseData.CandyType)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
