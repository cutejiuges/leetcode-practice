package LC2900

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/15 下午9:57
 * @FilePath: get_longest_subsequence_test.go
 * @Description:
 */

type TestData map[string]struct {
	Words     []string `yaml:"words"`
	Groups    []int    `yaml:"groups"`
	ExpectRes []string `yaml:"expect_res"`
}

func TestGetLongestSubsequence(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := getLongestSubsequence(caseData.Words, caseData.Groups); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
