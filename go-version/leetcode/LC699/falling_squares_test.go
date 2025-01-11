package LC699

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/7/28 上午8:29
 * @FilePath: falling_squares_test
 * @Description:
 */

type TestData map[string]struct {
	Positions [][]int `yaml:"Positions"`
	ExpectRes []int   `yaml:"ExpectRes"`
}

func TestFallingSquares(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := fallingSquares(caseData.Positions)
		if len(res) != len(caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectREs = %v", caseName, res, caseData.ExpectRes)
		}

		for i := range res {
			if res[i] != caseData.ExpectRes[i] {
				t.Errorf("%s, res = %v, ExpectREs = %v", caseName, res, caseData.ExpectRes)
			}
		}
	}
}
