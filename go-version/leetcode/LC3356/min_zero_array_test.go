package LC3356

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/22 上午12:19
 * @FilePath: min_zero_array_test.go
 * @Description:
 */

type TestData map[string]struct {
	Nums      []int   `yaml:"nums"`
	Queries   [][]int `yaml:"queries"`
	ExpectRes int     `yaml:"expect_res"`
}

func TestMinZeroArray(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := minZeroArray(caseData.Nums, caseData.Queries); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
