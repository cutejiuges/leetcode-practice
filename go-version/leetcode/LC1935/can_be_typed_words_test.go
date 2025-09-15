package LC1935

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/9/15 下午11:39
 * @FilePath: can_be_typed_words_test.go
 * @Description:
 */

type TestData map[string]struct {
	Text          string `yaml:"text"`
	BrokenLetters string `yaml:"broken_letters"`
	ExpectRes     int    `yaml:"expect_res"`
}

func TestCanBeTypedWords(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := canBeTypedWords(caseData.Text, caseData.BrokenLetters); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
