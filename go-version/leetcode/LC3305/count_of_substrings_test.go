package LC3305

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/9/30 下午10:01
 * @FilePath: count_of_substrings_test
 * @Description:
 */

type TestData map[string]struct {
	Word      string `yaml:"Word"`
	K         int    `yaml:"K"`
	ExpectRes int64  `yaml:"ExpectRes"`
}

func TestCountOfSubstrings(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := countOfSubstrings(caseData.Word, caseData.K)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
