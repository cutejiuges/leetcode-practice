package LC2273

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/10/13 下午11:00
 * @FilePath: remove_anagrams_test.go
 * @Description:
 */

type TestData map[string]struct {
	Words     []string `yaml:"words"`
	ExpectRes []string `yaml:"expect_res"`
}

func TestRemoveAnagrams(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		if res := removeAnagrams(caseData.Words); !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
