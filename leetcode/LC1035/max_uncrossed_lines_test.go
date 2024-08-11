package LC1035

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/8/11 上午11:36
 * @FilePath: max_uncrossed_lines_test
 * @Description:
 */

type TestData map[string]struct {
	Nums1     []int `yaml:"Nums1"`
	Nums2     []int `yaml:"Nums2"`
	ExpectRes int   `yaml:"ExpectRes"`
}

func TestMaxUncrossedLines(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := maxUncrossedLines(caseData.Nums1, caseData.Nums2)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectREs = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
