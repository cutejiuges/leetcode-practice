package LC2462

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/1 上午1:49
 * @FilePath: total_cost_test
 * @Description:
 */

type TestData map[string]struct {
	Costs      []int `yaml:"Costs"`
	K          int   `yaml:"K"`
	Candidates int   `yaml:"Candidates"`
	ExpectRes  int64 `yaml:"ExpectRes"`
}

func TestTotalCost(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for caseName, caseData := range testData {
		res := totalCost(caseData.Costs, caseData.K, caseData.Candidates)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
