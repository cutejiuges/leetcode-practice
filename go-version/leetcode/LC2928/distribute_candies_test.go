package LC2928

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/1 下午7:10
 * @FilePath: distribute_candies_test.yaml
 * @Description:
 */

type TestData map[string]struct {
	N         int `yaml:"N"`
	Limit     int `yaml:"Limit"`
	ExpectRes int `yaml:"ExpectRes"`
}

func TestDistributeCandies(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := distributeCandies(caseData.N, caseData.Limit)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
