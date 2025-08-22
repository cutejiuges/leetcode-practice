package LC3195

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/22 下午8:27
 * @FilePath: minimum_area_test.go
 * @Description:
 */

type TestData map[string]struct {
	Grid      [][]int `yaml:"grid"`
	ExpectRes int     `yaml:"expect_res"`
}

func TestMinimumArea(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := minimumArea(caseData.Grid); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
