package LC2269

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/10 下午11:29
 * @FilePath: divisor_substrings_test.go
 * @Description:
 */

type TestData map[string]struct {
	Num       int `yaml:"num"`
	K         int `yaml:"k"`
	ExpectRes int `yaml:"expect_res"`
}

func TestDivisorSubstrings(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := divisorSubstrings(caseData.Num, caseData.K)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
