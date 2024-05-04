package LC1652

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/5 上午12:32
 * @FilePath: decrypt_test
 * @Description:
 */

type TestData map[string]struct {
	Code      []int `yaml:"Code"`
	K         int   `yaml:"K"`
	ExpectRes []int `yaml:"ExpectRes"`
}

func TestDecrypt(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := decrypt(caseData.Code, caseData.K)
		if len(res) != len(caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectREs = %v", caseName, res, caseData.ExpectRes)
			for i := range res {
				if res[i] != caseData.ExpectRes[i] {
					t.Errorf("%s, res = %v, ExpectREs = %v", caseName, res, caseData.ExpectRes)
				}
			}
		}
	}
}
