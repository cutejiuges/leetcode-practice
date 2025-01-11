package LC3191

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/18 上午9:07
 * @FilePath: min_operations_test
 * @Description:
 */

type TestData map[string]struct {
	Nums      []int `yaml:"nums"`
	ExpectRes int   `yaml:"expect_res"`
}

func TestMinOperations(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := minOperations(caseData.Nums)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
