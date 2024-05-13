package LC994

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/13 下午10:58
 * @FilePath: oranges_rotting_test
 * @Description:
 */

type TestData map[string]struct {
	Grid      [][]int `yaml:"Grid"`
	ExpectRes int     `yaml:"ExpectRes"`
}

func TestOrangesRotting(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := orangesRotting(caseData.Grid)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
