package LC935

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/10 下午11:32
 * @FilePath: knight_dialer_test
 * @Description:
 */

type TestData map[string]struct {
	N         int `yaml:"n"`
	ExpectRes int `yaml:"expect_res"`
}

func TestKnightDialer(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := knightDialer(caseData.N)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
