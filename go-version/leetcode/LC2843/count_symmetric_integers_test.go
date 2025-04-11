package LC2843

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/11 下午9:59
 * @FilePath: count_symmetric_integers_test.go
 * @Description:
 */

type TestData map[string]struct {
	Low       int `yaml:"low"`
	High      int `yaml:"high"`
	ExpectRes int `yaml:"expect_res"`
}

func TestCountSymmetricIntegers(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := countSymmetricIntegers(caseData.Low, caseData.High)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
