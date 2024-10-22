package LC3184

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/22 下午11:13
 * @FilePath: count_complete_day_pairs_test
 * @Description:
 */

type TestData map[string]struct {
	Hours     []int `yaml:"hours"`
	ExpectRes int   `yaml:"expect_res"`
}

func TestCountCompleteDayPairs(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := countCompleteDayPairs(caseData.Hours)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
