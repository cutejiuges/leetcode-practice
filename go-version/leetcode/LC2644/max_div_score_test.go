package LC2644

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/22 上午3:23
 * @FilePath: max_div_score_test
 * @Description:
 */

type TestData map[string]struct {
	Nums      []int `yaml:"Nums"`
	Divisors  []int `yaml:"Divisors"`
	ExpectRes int   `yaml:"ExpectRes"`
}

func TestMaxDivScore(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := maxDivScore(caseData.Nums, caseData.Divisors)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
