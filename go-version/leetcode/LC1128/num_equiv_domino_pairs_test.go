package LC1128

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/4 上午12:44
 * @FilePath: num_equiv_domino_pairs_test.go
 * @Description:
 */

type TestData map[string]struct {
	Dominoes  [][]int `yaml:"dominoes"`
	ExpectRes int     `yaml:"expect_res"`
}

func TestNumEquivDominoPairs(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := numEquivDominoPairs(caseData.Dominoes); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
