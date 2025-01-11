package LC3222

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/5 下午11:40
 * @FilePath: losing_player_test
 * @Description:
 */

type TestData map[string]struct {
	X         int    `yaml:"x"`
	Y         int    `yaml:"y"`
	ExpectRes string `yaml:"expect_res"`
}

func TestLosingPlayer(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := losingPlayer(caseData.X, caseData.Y)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
