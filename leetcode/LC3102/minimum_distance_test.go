package LC3102

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/7/9 下午11:50
 * @FilePath: minimum_distance_test
 * @Description:
 */

type TestData map[string]struct {
	Points    [][]int `yaml:"Points"`
	ExpectRes int     `yaml:"ExpectRes"`
}

func TestMinimumDistance(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := minimumDistance(caseData.Points)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
