package LC3136

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/7/16 上午12:01
 * @FilePath: is_valid_test.go
 * @Description:
 */

type TestData map[string]struct {
	Word      string `yaml:"word"`
	ExpectRes bool   `yaml:"expect_res"`
}

func TestIsValid(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := isValid(caseData.Word); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
