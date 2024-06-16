package LC521

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/16 下午5:45
 * @FilePath: find_LUS_length_test
 * @Description:
 */

type TestData map[string]struct {
	A         string `yaml:"A"`
	B         string `yaml:"B"`
	ExpectRes int    `yaml:"ExpectRes"`
}

func TestFindLUSlength(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := findLUSlength(caseData.A, caseData.B)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
