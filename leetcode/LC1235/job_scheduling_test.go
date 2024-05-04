package LC1235

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/4 下午10:07
 * @FilePath: job_scheduling_test
 * @Description:
 */

type TestData map[string]struct {
	StartTime []int `yaml:"StartTime"`
	EndTime   []int `yaml:"EndTime"`
	Profit    []int `yaml:"Profit"`
	ExpectRes int   `yaml:"ExpectRes"`
}

func TestJobScheduling(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for caseName, caseData := range testData {
		res := jobScheduling(caseData.StartTime, caseData.EndTime, caseData.Profit)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
