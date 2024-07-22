package LC2101

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/7/22 下午11:26
 * @FilePath: maximum_detonation_test
 * @Description:
 */

type TestData map[string]struct {
	Bombs     [][]int `yaml:"Bombs"`
	ExpectRes int     `yaml:"ExpectRes"`
}

func TestMaximumDetonation(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := maximumDetonation(caseData.Bombs)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
