package LC743

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/25 下午11:04
 * @FilePath: network_delay_time_test
 * @Description:
 */

type TestData map[string]struct {
	Times     [][]int `yaml:"times"`
	N         int     `yaml:"n"`
	K         int     `yaml:"k"`
	ExpectRes int     `yaml:"expect_res"`
}

func TestNetworkDelayTime(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := networkDelayTime(caseData.Times, caseData.N, caseData.K)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
