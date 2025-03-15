package LC3110

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/15 下午9:33
 * @FilePath: score_of_string_test
 * @Description:
 */

type TestData map[string]struct {
	S         string `yaml:"s"`
	ExpectRes int    `yaml:"expect_res"`
}

func TestScoreOfString(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := scoreOfString(caseData.S)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
