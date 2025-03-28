package LC2716

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/28 下午11:25
 * @FilePath: minimized_string_length_test.go
 * @Description:
 */

type TestData map[string]struct {
	S         string `yaml:"s"`
	ExpectRes int    `yaml:"expect_res"`
}

func TestMinimizedStringLength(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := minimizedStringLength(caseData.S)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
