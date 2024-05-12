package LC1553

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/12 下午8:00
 * @FilePath: min_days_test
 * @Description:
 */

type TestData map[string]struct {
	N         int `yaml:"N"`
	ExpectRes int `yaml:"ExpectRes"`
}

func TestMinDays(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := minDays(caseData.N)
		if res != caseData.ExpectRes {
			t.Errorf("minDays # %s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}

		res = minDays2(caseData.N)
		if res != caseData.ExpectRes {
			t.Errorf("minDays2 # %s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
