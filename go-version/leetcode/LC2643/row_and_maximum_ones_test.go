package LC2643

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/22 下午1:04
 * @FilePath: row_and_maximum_ones_test
 * @Description:
 */

type TestData map[string]struct {
	Mat       [][]int `yaml:"mat"`
	ExpectRes []int   `yaml:"expect_res"`
}

func TestRowAndMaximumOnes(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := rowAndMaximumOnes(caseData.Mat)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectREs = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
