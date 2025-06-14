package LC2566

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/6/14 下午9:00
 * @FilePath: min_max_difference_test.go
 * @Description:
 */

type TestData map[string]struct {
	Num       int `yaml:"num"`
	ExpectRes int `yaml:"expect_res"`
}

func TestMinMaxDifference(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := minMaxDifference(caseData.Num); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
