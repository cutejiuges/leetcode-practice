package LC2434

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/6/6 下午10:39
 * @FilePath: robot_with_string_test.go
 * @Description:
 */

type TestData map[string]struct {
	S         string `yaml:"s"`
	ExpectRes string `yaml:"expect_res"`
}

func TestRobotWithString(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := robotWithString(caseData.S); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
