package LC3143

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/8/3 下午9:47
 * @FilePath: max_points_inside_square_test
 * @Description:
 */

type TestData map[string]struct {
	Points    [][]int `yaml:"Points"`
	S         string  `yaml:"S"`
	ExpectRes int     `yaml:"ExpectRes"`
}

func TestMaxPointsInsideSquare(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := maxPointsInsideSquare(caseData.Points, caseData.S)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
