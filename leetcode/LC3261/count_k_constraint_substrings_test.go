package LC3261

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/13 下午9:49
 * @FilePath: count_k_constraint_substrings_test
 * @Description:
 */

type TestData map[string]struct {
	S         string  `yaml:"s"`
	K         int     `yaml:"k"`
	Queries   [][]int `yaml:"queries"`
	ExpectRes []int64 `yaml:"expect_res"`
}

func TestCountKConstraintSubstrings(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := countKConstraintSubstrings(caseData.S, caseData.K, caseData.Queries)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
