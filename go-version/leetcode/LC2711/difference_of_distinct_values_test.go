package LC2711

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/25 下午11:15
 * @FilePath: difference_of_distinct_values_test.go
 * @Description:
 */

type TestData map[string]struct {
	Grid      [][]int `yaml:"grid"`
	ExpectRes [][]int `yaml:"expect_res"`
}

func TestDifferenceOfDistinctValues(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := differenceOfDistinctValues(caseData.Grid)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
