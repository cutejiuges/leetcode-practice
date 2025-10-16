package LC2598

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/10/16 下午11:41
 * @FilePath: find_smallest_integer_test.go
 * @Description:
 */

type TestData map[string]struct {
	Nums      []int `yaml:"nums"`
	Value     int   `yaml:"value"`
	ExpectRes int   `yaml:"expect_res"`
}

func TestFindSmallestInteger(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := findSmallestInteger(caseData.Nums, caseData.Value); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectREs = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
