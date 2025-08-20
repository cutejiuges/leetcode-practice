package LC1277

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/20 上午8:43
 * @FilePath: count_square_test.yaml.go
 * @Description:
 */

type TestData map[string]struct {
	Matrix    [][]int `yaml:"matrix"`
	ExpectRes int     `yaml:"expect_res"`
}

func TestCountSquares(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := countSquares(caseData.Matrix); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, expect %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
