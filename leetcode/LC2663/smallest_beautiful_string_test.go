package LC2663

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/22 上午3:08
 * @FilePath: smallest_beautiful_string_test
 * @Description:
 */

type TestData map[string]struct {
	S         string `yaml:"S"`
	K         int    `yaml:"K"`
	ExpectRes string `yaml:"ExpectRes"`
}

func TestSmallestBeautifulString(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := smallestBeautifulString(caseData.S, caseData.K)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
