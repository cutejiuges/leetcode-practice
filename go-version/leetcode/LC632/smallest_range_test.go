package LC632

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/24 下午12:44
 * @FilePath: smallest_range_test
 * @Description:
 */

type TestData map[string]struct {
	Nums      [][]int `yaml:"nums"`
	ExpectRes []int   `yaml:"expect_res"`
}

func TestSmallestRange(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := smallestRange(caseData.Nums)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
