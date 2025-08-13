package LC326

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/13 下午11:22
 * @FilePath: is_power_of_three_test.go
 * @Description:
 */

type TestData map[string]struct {
	N         int  `yaml:"n"`
	ExpectRes bool `yaml:"expect_res"`
}

func TestIsPowerOfThree(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := isPowerOfThree(caseData.N); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
