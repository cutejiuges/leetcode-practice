package LC2712

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/27 下午11:39
 * @FilePath: minimum_cost_test.go
 * @Description:
 */

type TestData map[string]struct {
	S         string `yaml:"s"`
	ExpectRes int64  `yaml:"expect_res"`
}

func TestMinimumCost(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := minimumCost(caseData.S)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
