package LC1578

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/11/3 下午11:27
 * @FilePath: min_cost_test.go
 * @Description:
 */

type TestData map[string]struct {
	Colors     string `yaml:"colors"`
	NeededTime []int  `yaml:"needed_time"`
	ExpectRes  int    `yaml:"expect_res"`
}

func TestMinCost(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := minCost(caseData.Colors, caseData.NeededTime); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectREs = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
