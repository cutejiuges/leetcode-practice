package LC2094

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/12 上午9:23
 * @FilePath: find_even_numbers_test.go
 * @Description:
 */

type TestData map[string]struct {
	Digits    []int `yaml:"digits"`
	ExpectRes []int `yaml:"expect_res"`
}

func TestFindEvenNumbers(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := findEvenNumbers(caseData.Digits); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
