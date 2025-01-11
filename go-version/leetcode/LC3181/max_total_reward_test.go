package LC3181

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/26 下午6:56
 * @FilePath: max_total_reward_test
 * @Description:
 */

type TestData map[string]struct {
	RewardValues []int `yaml:"reward_values"`
	ExpectRes    int   `yaml:"expect_res"`
}

func TestMaxTotalReward(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := maxTotalReward(caseData.RewardValues)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
