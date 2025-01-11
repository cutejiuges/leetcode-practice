package LC3297

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/1/9 下午11:39
 * @FilePath: valid_substring_count_test
 * @Description:
 */

type TestData map[string]struct {
	Word1     string `yaml:"word1"`
	Word2     string `yaml:"word2"`
	ExpectRes int64  `yaml:"expect_res"`
}

func TestValidSubstringCount(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := validSubstringCount(caseData.Word1, caseData.Word2)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
