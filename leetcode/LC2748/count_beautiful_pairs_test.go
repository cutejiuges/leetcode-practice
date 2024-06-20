package LC2748

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/20 下午11:01
 * @FilePath: count_beautiful_pairs_test
 * @Description:
 */

type TestData map[string]struct {
	Nums      []int `yaml:"Nums"`
	ExpectRes int   `yaml:"ExpectRes"`
}

func TestCountBeautifulPairs(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := countBeautifulPairs(caseData.Nums)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
