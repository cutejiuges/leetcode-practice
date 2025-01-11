package LC1103

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/3 下午11:03
 * @FilePath: distribute_candies_test
 * @Description:
 */

type TestData map[string]struct {
	Candies   int   `yaml:"Candies"`
	NumPeople int   `yaml:"NumPeople"`
	ExpectRes []int `yaml:"ExpectRes"`
}

func TestDistributeCandies(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := distributeCandies(caseData.Candies, caseData.NumPeople)
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
