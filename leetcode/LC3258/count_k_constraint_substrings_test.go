package LC3258

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/12 下午11:11
 * @FilePath: count_k_constraint_substrings_test
 * @Description:
 */

type TestData map[string]struct {
	S         string `yaml:"s"`
	K         int    `yaml:"k"`
	ExpectRes int    `yaml:"expect_res"`
}

func TestCountKConstraintSubstrings(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := countKConstraintSubstrings(caseData.S, caseData.K)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
