package LC2244

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/14 下午11:33
 * @FilePath: minimum_rounds_test
 * @Description:
 */

type TestData map[string]struct {
	Tasks     []int `yaml:"Tasks"`
	ExpectRes int   `yaml:"ExpectRes"`
}

func TestMinimumRounds(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := minimumRounds(caseData.Tasks)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
