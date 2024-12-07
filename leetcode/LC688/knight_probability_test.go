package LC688

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/7 上午8:04
 * @FilePath: knight_probability_test
 * @Description:
 */

type TestData map[string]struct {
	N         int     `yaml:"n"`
	K         int     `yaml:"k"`
	Row       int     `yaml:"row"`
	Column    int     `yaml:"column"`
	ExpectRes float64 `yaml:"expect_res"`
}

func TestKnightProbability(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := knightProbability(caseData.N, caseData.K, caseData.Row, caseData.Column)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
