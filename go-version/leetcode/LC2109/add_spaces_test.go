package LC2109

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/30 上午8:44
 * @FilePath: add_spaces_test.go
 * @Description:
 */

type TestData map[string]struct {
	S         string `yaml:"s"`
	Spaces    []int  `yaml:"spaces"`
	ExpectRes string `yaml:"expect_res"`
}

func TestAddSpaces(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := addSpaces(caseData.S, caseData.Spaces)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
