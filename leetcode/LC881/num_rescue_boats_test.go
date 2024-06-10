package LC881

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/10 下午11:58
 * @FilePath: num_rescue_boats_test
 * @Description:
 */

type TestData map[string]struct {
	People    []int `yaml:"People"`
	Limit     int   `yaml:"Limit"`
	ExpectRes int   `yaml:"ExpectRes"`
}

func TestNumRescueBoats(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := numRescueBoats(caseData.People, caseData.Limit)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v\n", caseName, res, caseData.ExpectRes)
		}
	}
}
