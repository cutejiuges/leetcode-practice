package LC66

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2026/1/1 下午10:19
 * @FilePath: plus_one_test.go
 * @Description:
 */

type TestData map[string]struct {
	Digits    []int `yaml:"digits"`
	ExpectRes []int `yaml:"expect_res"`
}

func TestPlusOne(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			ans := plusOne(caseData.Digits)
			if !reflect.DeepEqual(ans, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRTes = %v", caseName, ans, caseData.ExpectRes)
			}
		})
	}
}
