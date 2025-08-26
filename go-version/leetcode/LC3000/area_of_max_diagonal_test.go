package LC3000

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/26 上午8:21
 * @FilePath: area_of_max_diagonal_test.go
 * @Description:
 */

type TestData map[string]struct {
	Dimensions [][]int `yaml:"dimensions"`
	ExpectRes  int     `yaml:"expect_res"`
}

func TestAreaOfMaxDiagonal(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := areaOfMaxDiagonal(caseData.Dimensions); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
