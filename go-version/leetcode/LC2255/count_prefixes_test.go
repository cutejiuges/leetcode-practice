package LC2255

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/24 下午10:35
 * @FilePath: count_prefixes_test
 * @Description:
 */

type TestData map[string]struct {
	Words     []string `yaml:"words"`
	S         string   `yaml:"s"`
	ExpectRes int      `yaml:"expect_res"`
}

func TestCountPrefixes(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := countPrefixes(caseData.Words, caseData.S)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
