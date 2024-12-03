package LC3274

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/3 下午10:36
 * @FilePath: check_two_chessboards_test
 * @Description:
 */

type TestData map[string]struct {
	coordinate1 string `yaml:"coordinate1"`
	coordinate2 string `yaml:"coordinate2"`
	expectRes   bool   `yaml:"expect_res"`
}

func TestCheckTwoChessboards(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := checkTwoChessboards(caseData.coordinate1, caseData.coordinate2)
		if !reflect.DeepEqual(res, caseData.expectRes) {
			t.Errorf("%s, res = %v, expectRes = %v", caseName, res, caseData.expectRes)
		}
	}
}
