package LC3238

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/23 下午11:30
 * @FilePath: winning_player_count_test
 * @Description:
 */

type TestData map[string]struct {
	N         int     `yaml:"n"`
	Pick      [][]int `yaml:"pick"`
	ExpectRes int     `yaml:"expect_res"`
}

func TestWinningPlayerCount(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := winningPlayerCount(caseData.N, caseData.Pick)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
