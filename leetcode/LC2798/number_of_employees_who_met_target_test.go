package LC2798

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/4/30 上午12:13
 * @FilePath: number_of_employees_who_met_target_test
 * @Description:
 */

type TestData map[string]struct {
	Hours     []int `yaml:"Hours"`
	Target    int   `yaml:"Target"`
	ExpectRes int   `yaml:"ExpectRes"`
}

func TestNumberOfEmployeesWhoMetTarget(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for caseName, caseData := range testData {
		res := numberOfEmployeesWhoMetTarget(caseData.Hours, caseData.Target)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
