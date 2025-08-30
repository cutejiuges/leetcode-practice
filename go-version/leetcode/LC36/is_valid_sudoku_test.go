package LC36

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/30 上午11:35
 * @FilePath: is_valid_sudoku_test.go
 * @Description:
 */

type TestData map[string]struct {
	Board     [][]string `yaml:"board"`
	ExpectRes bool       `yaml:"expect_res"`
}

func TestIsValidSudoku(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			board := make([][]byte, len(caseData.Board))
			for i, row := range caseData.Board {
				board[i] = make([]byte, len(row))
				for j, cell := range row {
					board[i][j] = cell[0]
				}
			}
			if res := isValidSudoku(board); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
