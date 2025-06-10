package LC3442

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/6/10 下午10:58
 * @FilePath: max_difference_test.go
 * @Description:
 */

type TestData map[string]struct {
	S         string `yaml:"s"`
	ExpectRes int    `yaml:"expect_res"`
}

func TestMaxDifference(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := maxDifference(caseData.S); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
