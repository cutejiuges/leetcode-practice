package LC3304

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/9/30 下午9:31
 * @FilePath: kth_character_test
 * @Description:
 */

type TestData map[string]struct {
	K         int    `yaml:"K"`
	ExpectRes string `yaml:"ExpectRes"`
}

func TestKthCharacter(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := kthCharacter(caseData.K)
		if res != caseData.ExpectRes[0] {
			t.Errorf("%s, res = %c, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
