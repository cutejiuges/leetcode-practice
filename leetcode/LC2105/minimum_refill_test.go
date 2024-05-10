package LC2105

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/11 上午12:05
 * @FilePath: minimum_refill_test
 * @Description:
 */

type TestData map[string]struct {
	Plants    []int `yaml:"Plants"`
	CapacityA int   `yaml:"CapacityA"`
	CapacityB int   `yaml:"CapacityB"`
	ExpectRes int   `yaml:"ExpectRes"`
}

func TestMinimumRefill(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := minimumRefill(caseData.Plants, caseData.CapacityA, caseData.CapacityB)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
