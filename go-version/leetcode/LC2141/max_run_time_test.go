package LC2141

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/12/1 上午9:16
 * @FilePath: max_run_time_test.go
 * @Description:
 */

type TestData map[string]struct {
	N         int   `yaml:"n"`
	Batteries []int `yaml:"batteries"`
	ExpectRes int64 `yaml:"expect_res"`
}

func TestMaxRunTime(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := maxRunTime(caseData.N, caseData.Batteries); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
