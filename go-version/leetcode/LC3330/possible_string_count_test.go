package LC3330

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/7/1 上午8:49
 * @FilePath: possible_string_count_test.go
 * @Description:
 */

type TestData map[string]struct {
	Word      string `yaml:"word"`
	ExpectRes int    `yaml:"expect_res"`
}

func TestPossibleStringCount(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := possibleStringCount(caseData.Word); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
