package LC2732

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/25 下午10:41
 * @FilePath: good_subset_of_binary_matrix_test
 * @Description:
 */
type TestData map[string]struct {
	Grid      [][]int `yaml:"Grid"`
	ExpectRes []int   `yaml:"ExpectRes"`
}

func TestGoodSubsetofBinaryMatrix(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := goodSubsetofBinaryMatrix(caseData.Grid)
		if len(res) != len(caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
		for i := range res {
			if res[i] != caseData.ExpectRes[i] {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		}
	}
}
