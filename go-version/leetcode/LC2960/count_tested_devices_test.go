package LC2960

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/10 下午11:47
 * @FilePath: count_tested_devices_test
 * @Description:
 */

type TestData map[string]struct {
	BatteryPercentages []int `yaml:"BatteryPercentages"`
	ExpectRes          int   `yaml:"ExpectRes"`
}

func TestCountTestedDevices(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := countTestedDevices(caseData.BatteryPercentages)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
