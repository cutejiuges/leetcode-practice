package LC1535

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/29 下午9:14
 * @FilePath: get_winner_test
 * @Description:
 */

type TestData map[string]struct {
	Arr       []int `yaml:"Arr"`
	K         int   `yaml:"K"`
	ExpectRes int   `yaml:"ExpectRes"`
}

func TestGetWinner(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := getWinner(caseData.Arr, caseData.K)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
