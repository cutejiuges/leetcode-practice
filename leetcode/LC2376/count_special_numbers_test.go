package LC2376

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/9/20 下午11:29
 * @FilePath: count_special_numbers_test
 * @Description:
 */

type TestData map[string]struct {
	N         int `yaml:"N"`
	ExpectRes int `yaml:"ExpectRes"`
}

func TestCountSpecialNumbers(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := countSpecialNumbers(caseData.N)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
