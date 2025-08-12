package LC2787

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/12 下午11:09
 * @FilePath: number_of_ways_test.go
 * @Description:
 */

type TestData map[string]struct {
	N         int `yaml:"n"`
	X         int `yaml:"x"`
	ExpectRes int `yaml:"expect_res"`
}

func TestNumberOfWays(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := numberOfWays(caseData.N, caseData.X); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
