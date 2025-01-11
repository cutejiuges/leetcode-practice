package LC3011

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/7/13 上午10:04
 * @FilePath: can_sort_array_test
 * @Description:
 */

type TestData map[string]struct {
	Nums      []int `yaml:"Nums"`
	ExpectRes bool  `yaml:"ExpectRes"`
}

func TestCanSortArray(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := canSortArray(caseData.Nums)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectREs = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
