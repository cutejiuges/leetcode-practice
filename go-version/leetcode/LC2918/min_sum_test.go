package LC2918

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/10 下午8:44
 * @FilePath: min_sum_test.go
 * @Description:
 */

type TestData map[string]struct {
	Nums1     []int `yaml:"nums1"`
	Nums2     []int `yaml:"nums2"`
	ExpectRes int64 `yaml:"expect_res"`
}

func TestMinSun(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := minSum(caseData.Nums1, caseData.Nums2); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
