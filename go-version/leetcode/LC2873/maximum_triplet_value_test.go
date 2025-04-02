package LC2873

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/2 下午11:07
 * @FilePath: maximum_triplet_value_test.go
 * @Description:
 */

type TestData map[string]struct {
	Nums      []int `yaml:"nums"`
	ExpectRes int64 `yaml:"expect_res"`
}

func TestMaximumTripletValue(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := maximumTripletValue(caseData.Nums)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
