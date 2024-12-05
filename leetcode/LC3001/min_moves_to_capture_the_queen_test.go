package LC3001

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/5 下午11:03
 * @FilePath: min_moves_to_capture_the_queen_test
 * @Description:
 */

type TestData map[string]struct {
	A         int `yaml:"a"`
	B         int `yaml:"b"`
	C         int `yaml:"c"`
	D         int `yaml:"d"`
	E         int `yaml:"e"`
	F         int `yaml:"f"`
	ExpectRes int `yaml:"expect_res"`
}

func TestMinMovesToCaptureTheQueen(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := minMovesToCaptureTheQueen(caseData.A, caseData.B, caseData.C, caseData.D, caseData.E, caseData.F)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, expectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
