package LC3146

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/8/24 上午1:10
 * @FilePath: find_permutation_difference_test
 * @Description:
 */

type TestData map[string]struct {
	S         string `yaml:"S"`
	T         string `yaml:"T"`
	ExpectRes int    `yaml:"ExpectRes"`
}

func TestFindPermutationDifference(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := findPermutationDifference(caseData.S, caseData.T)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, expectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
