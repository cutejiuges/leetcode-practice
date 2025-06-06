package LC2302

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/28 下午11:32
 * @FilePath: count_subarrays_test.go
 * @Description:
 */

type TestData map[string]struct {
	Nums      []int `yaml:"nums"`
	K         int64 `yaml:"k"`
	ExpectRes int64 `yaml:"expect_res"`
}

func TestCountSubarrays(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := countSubarrays(caseData.Nums, caseData.K); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
