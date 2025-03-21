package LC2680

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/21 下午10:40
 * @FilePath: maximum_or_test
 * @Description:
 */

type TestData map[string]struct {
	Nums      []int `yaml:"nums"`
	K         int   `yaml:"k"`
	ExpectRes int64 `yaml:"expect_res"`
}

func TestMaximumOr(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := maximumOr(caseData.Nums, caseData.K)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
