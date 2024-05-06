package LC741

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/7 上午12:06
 * @FilePath: cherry_pickup_test
 * @Description:
 */

type TestData map[string]struct {
	Grid      [][]int `yaml:"Grid"`
	ExpectRes int     `yaml:"ExpectRes"`
}

func TestCherryPickup(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := cherryPickup(caseData.Grid)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
