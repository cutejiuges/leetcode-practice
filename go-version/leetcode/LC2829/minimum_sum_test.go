package LC2829

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/26 上午8:37
 * @FilePath: minimum_sum_test.go
 * @Description:
 */

type TestData map[string]struct {
	N         int `yaml:"n"`
	K         int `yaml:"k"`
	ExpectRes int `yaml:"expect_res"`
}

func TestMinimumSum(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := minimumSum(caseData.N, caseData.K)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
