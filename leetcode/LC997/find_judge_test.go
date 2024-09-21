package LC997

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/9/22 上午12:22
 * @FilePath: find_judge_test
 * @Description:
 */

type TestData map[string]struct {
	N         int     `yaml:"N"`
	Trust     [][]int `yaml:"Trust"`
	ExpectRes int     `yaml:"ExpectRes"`
}

func TestFindJudge(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := findJudge(caseData.N, caseData.Trust)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
