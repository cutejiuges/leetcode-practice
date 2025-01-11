package LC3200

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/15 下午11:20
 * @FilePath: max_height_of_triangle_test
 * @Description:
 */

type TestData map[string]struct {
	Red       int `yaml:"red"`
	Blue      int `yaml:"blue"`
	ExpectRes int `yaml:"expect_res"`
}

func TestMaxHeightOfTriangle(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := maxHeightOfTriangle(caseData.Red, caseData.Blue)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
