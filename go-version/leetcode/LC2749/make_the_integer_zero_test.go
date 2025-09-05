package LC2749

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/9/5 下午11:47
 * @FilePath: make_the_integer_zero_test.go
 * @Description:
 */

type TestData map[string]struct {
	Num1      int `yaml:"num1"`
	Num2      int `yaml:"num2"`
	ExpectRes int `yaml:"expect_res"`
}

func TestMakeTheIntegerZero(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := makeTheIntegerZero(caseData.Num1, caseData.Num2); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
