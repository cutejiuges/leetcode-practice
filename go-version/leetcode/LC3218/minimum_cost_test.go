package LC3218

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/25 下午11:09
 * @FilePath: minimum_cost_test
 * @Description:
 */

type TestData map[string]struct {
	M             int   `yaml:"m"`
	N             int   `yaml:"n"`
	HorizontalCut []int `yaml:"horizontal_cut"`
	VerticalCut   []int `yaml:"vertical_cut"`
	ExpectRes     int   `yaml:"expect_res"`
}

func TestMinimumCost(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := minimumCost(caseData.M, caseData.N, caseData.HorizontalCut, caseData.VerticalCut)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}