package LC2894

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/27 下午11:24
 * @FilePath: difference_of_sums_test.go
 * @Description:
 */

type TestData map[string]struct {
	N         int `yaml:"n"`
	M         int `yaml:"m"`
	ExpectRes int `yaml:"expect_res"`
}

func TestDifferenceOfSums(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := differenceOfSums(caseData.N, caseData.M); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
