package LC3194

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/16 上午12:48
 * @FilePath: minimum_average_test
 * @Description:
 */

type TestData map[string]struct {
	Nums      []int   `yaml:"nums"`
	ExpectRes float64 `yaml:"expect_res"`
}

func TestMinimumAverage(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := minimumAverage(caseData.Nums)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
