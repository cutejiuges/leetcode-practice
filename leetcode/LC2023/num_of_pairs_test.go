package LC2023

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/23 下午9:25
 * @FilePath: num_of_pairs_test
 * @Description:
 */

type TestData map[string]struct {
	Nums      []string `yaml:"Nums"`
	Target    string   `yaml:"Target"`
	ExpectRes int      `yaml:"ExpectRes"`
}

func TestNumOfPairs(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := numOfPairs(caseData.Nums, caseData.Target)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
