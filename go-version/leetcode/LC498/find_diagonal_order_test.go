package LC498

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/25 上午7:26
 * @FilePath: find_diagonal_order_test.go
 * @Description:
 */

type TestData map[string]struct {
	Mat       [][]int `yaml:"mat"`
	ExpectRes []int   `yaml:"expect_res"`
}

func TestFindDiagonalOrder(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := findDiagonalOrder(caseData.Mat); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
