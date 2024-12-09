package LC1812

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/9 下午10:38
 * @FilePath: square_is_white_test
 * @Description:
 */

type TestData map[string]struct {
	Coordinates string `yaml:"coordinates"`
	ExpectRes   bool   `yaml:"expect_res"`
}

func TestSquareIsWhite(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := squareIsWhite(caseData.Coordinates)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
