package LC2145

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/21 上午9:02
 * @FilePath: number_of_arrays_test.go
 * @Description:
 */

type TestData map[string]struct {
	Differences []int `yaml:"differences"`
	Upper       int   `yaml:"upper"`
	Lower       int   `yaml:"lower"`
	ExpectRes   int   `yaml:"expect_res"`
}

func TestNumberOfArrays(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := numberOfArrays(caseData.Differences, caseData.Lower, caseData.Upper); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
