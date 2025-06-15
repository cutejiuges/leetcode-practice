package LC1432

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/6/15 下午8:21
 * @FilePath: max_diff_test.go
 * @Description:
 */

type TestData map[string]struct {
	Num       int `yaml:"num"`
	ExpectRes int `yaml:"expect_res"`
}

func TestMaxDiff(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := maxDiff(caseData.Num); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
