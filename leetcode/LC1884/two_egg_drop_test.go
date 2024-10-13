package LC1884

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/13 下午5:52
 * @FilePath: two_egg_drop_test
 * @Description:
 */

type TestData map[string]struct {
	N         int `yaml:"n"`
	ExpectRes int `yaml:"expect_res"`
}

func TestTwoEggDrop(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := twoEggDrop(caseData.N)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
