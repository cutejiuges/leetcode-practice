package LC807

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/7/14 下午3:44
 * @FilePath: max_increase_keeping_skyline_test
 * @Description:
 */

type TestData map[string]struct {
	Grid      [][]int `yaml:"Grid"`
	ExpectRes int     `yaml:"ExpectRes"`
}

func TestMaxIncreaseKeepingSkyline(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := maxIncreaseKeepingSkyline(caseData.Grid)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
