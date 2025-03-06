package LC2588

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/6 下午11:39
 * @FilePath: beautiful_subarrays_test
 * @Description:
 */

type TestData map[string]struct {
	Nums      []int `yaml:"nums"`
	ExpectRes int64 `yaml:"expect_res"`
}

func TestBeautifulSubarrays(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := beautifulSubarrays(caseData.Nums)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData)
			}
		})
	}
}
