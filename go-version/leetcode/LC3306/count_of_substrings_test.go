package LC3306

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/13 下午11:39
 * @FilePath: count_of_substrings_test
 * @Description:
 */

type TestData map[string]struct {
	Word      string `yaml:"word"`
	K         int    `yaml:"k"`
	ExpectRes int64  `yaml:"expect_res"`
}

func TestCountOfSubstrings(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := countOfSubstrings(caseData.Word, caseData.K)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
