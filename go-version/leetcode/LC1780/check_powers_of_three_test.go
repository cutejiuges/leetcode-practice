package LC1780

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/14 下午11:01
 * @FilePath: check_powers_of_three_test.go
 * @Description:
 */

type TestData map[string]struct {
	N         int  `yaml:"n"`
	ExpectRes bool `yaml:"expect_res"`
}

func TestCheckPowersOfThree(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := checkPowersOfThree(caseData.N); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
