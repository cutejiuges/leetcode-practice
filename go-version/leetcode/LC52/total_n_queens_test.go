package LC52

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/2 下午11:17
 * @FilePath: total_n_queens_test
 * @Description:
 */

type TestData map[string]struct {
	n         int `yaml:"n"`
	expectRes int `yaml:"expect_res"`
}

func TestTotalNQueues(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := totalNQueens(caseData.n)
		if !reflect.DeepEqual(res, caseData.expectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.expectRes)
		}
	}
}
