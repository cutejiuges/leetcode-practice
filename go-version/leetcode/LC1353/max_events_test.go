package LC1353

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/7/7 上午9:11
 * @FilePath: max_events_test.go
 * @Description:
 */

type TestData map[string]struct {
	Events    [][]int `yaml:"events"`
	ExpectRes int     `yaml:"expect_res"`
}

func TestMaxEvents(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := maxEvents(caseData.Events); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
