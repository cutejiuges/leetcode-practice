package LC1922

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/13 下午6:53
 * @FilePath: count_good_numbers_test.go
 * @Description:
 */

type TestData map[string]struct {
	N         int64 `yaml:"n"`
	ExpectRes int   `yaml:"expect_res"`
}

func TestCountGoodNumbers(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := countGoodNumbers(caseData.N)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
