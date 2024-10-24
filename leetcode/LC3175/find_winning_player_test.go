package LC3175

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/24 下午11:29
 * @FilePath: find_winning_player_test
 * @Description:
 */

type TestData map[string]struct {
	Skills    []int `yaml:"skills"`
	K         int   `yaml:"k"`
	ExpectRes int   `yaml:"expect_res"`
}

func TestFindWinningPlayer(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := findWinningPlayer(caseData.Skills, caseData.K)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
