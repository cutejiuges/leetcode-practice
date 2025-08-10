package LC869

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/10 下午10:59
 * @FilePath: reordered_power_of_2_test.go
 * @Description:
 */

type TestData map[string]struct {
	N         int  `yaml:"n"`
	ExpectRes bool `yaml:"expect_res"`
}

func TestReorderedPowerOf2(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := reorderedPowerOf2(caseData.N); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectREs = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
