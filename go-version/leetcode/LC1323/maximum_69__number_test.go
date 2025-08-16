package LC1323

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/16 下午10:07
 * @FilePath: maximum_69__number_test.go
 * @Description:
 */

type TestData map[string]struct {
	Num       int `yaml:"num"`
	ExpectRes int `yaml:"expect_res"`
}

func TestMaximum69Number(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := maximum69Number(caseData.Num); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
