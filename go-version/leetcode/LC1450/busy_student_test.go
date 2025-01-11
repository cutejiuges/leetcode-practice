package LC1450

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/9/1 上午9:14
 * @FilePath: busy_student_test
 * @Description:
 */

type TestData map[string]struct {
	StartTime []int `yaml:"StartTime"`
	EndTime   []int `yaml:"EndTime"`
	QueryTime int   `yaml:"QueryTime"`
	ExpectRes int   `yaml:"ExpectRes"`
}

func TestBusyStudent(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := busyStudent(caseData.StartTime, caseData.EndTime, caseData.QueryTime)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
