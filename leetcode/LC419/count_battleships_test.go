package LC419

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/11 上午12:16
 * @FilePath: count_battleships_test
 * @Description:
 */

type TestData map[string]struct {
	Board     [][]byte `yaml:"Board"`
	ExpectRes int      `yaml:"ExpectRes"`
}

func TestCountBattleships(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := countBattleships(caseData.Board)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
