package LC838

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/2 上午1:09
 * @FilePath: push_dominoes_test.go
 * @Description:
 */

type TestData map[string]struct {
	Dominoes  string `yaml:"dominoes"`
	ExpectRes string `yaml:"expect_res"`
}

func TestPushDominoes(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := pushDominoes(caseData.Dominoes); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
