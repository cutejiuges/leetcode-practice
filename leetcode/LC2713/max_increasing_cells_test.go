package LC2713

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/19 下午11:21
 * @FilePath: max_increasing_cells_test
 * @Description:
 */

type TestData map[string]struct {
	Mat       [][]int `yaml:"Mat"`
	ExpectRes int     `yaml:"ExpectRes"`
}

func TestMaxIncreasingCells(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := maxIncreasingCells(caseData.Mat)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
