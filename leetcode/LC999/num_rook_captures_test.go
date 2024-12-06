package LC999

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/6 下午10:57
 * @FilePath: num_rook_captures_test
 * @Description:
 */

type TestData map[string]struct {
	Board     [][]string `yaml:"board"`
	ExpectRes int        `yaml:"expect_res"`
}

func TestNumRookCaptures(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		board := make([][]byte, len(caseData.Board))
		for i, row := range caseData.Board {
			for _, ch := range row {
				board[i] = append(board[i], ch[0])
			}
		}

		res := numRookCaptures(board)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
