package LC120

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/9/26 上午12:07
 * @FilePath: minimum_total_test.go
 * @Description:
 */

type TestData map[string]struct {
	Triangle  [][]int `yaml:"triangle"`
	ExpectRes int     `yaml:"expect_res"`
}

func TestMinimumTotal(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := minimumTotal(caseData.Triangle); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
