package LC2942

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/24 下午4:52
 * @FilePath: find_words_containing_test.go
 * @Description:
 */

type TestData map[string]struct {
	Words     []string `yaml:"words"`
	X         string   `yaml:"x"`
	ExpectRes []int    `yaml:"expect_res"`
}

func TestFindWordsContaining(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := findWordsContaining(caseData.Words, caseData.X[0]); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
