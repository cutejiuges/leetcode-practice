package LC812

import (
	"cutejiuge/leetcode-practice/util"
	"math"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/9/27 下午10:51
 * @FilePath: largest_triangle_area_test.go
 * @Description:
 */

type TestData map[string]struct {
	Points    [][]int `yaml:"points"`
	ExpectRes float64 `yaml:"expect_res"`
}

func TestLargestTriangleArea(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := largestTriangleArea(caseData.Points); math.Abs(res-caseData.ExpectRes) > 1e-5 {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
