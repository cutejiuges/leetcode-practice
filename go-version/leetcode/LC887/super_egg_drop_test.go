package LC887

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/14 下午11:44
 * @FilePath: super_egg_drop_test
 * @Description:
 */

type TestData map[string]struct {
	K         int `yaml:"k"`
	N         int `yaml:"n"`
	ExpectRes int `yaml:"expect_res"`
}

func TestSuperEggDrop(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := superEggDrop(caseData.K, caseData.N)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
