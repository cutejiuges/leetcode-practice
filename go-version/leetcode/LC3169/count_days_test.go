package LC3169

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/7/11 上午9:09
 * @FilePath: count_days_test.go
 * @Description:
 */

type TestData map[string]struct {
	Days      int     `yaml:"days"`
	Meetings  [][]int `yaml:"meetings"`
	ExpectRes int     `yaml:"expect_res"`
}

func TestCountDays(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := countDays(caseData.Days, caseData.Meetings); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
