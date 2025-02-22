package LC2506

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/2/22 下午9:44
 * @FilePath: similar_pairs_test
 * @Description: 测试功能
 */

type TestData map[string]struct {
	Words     []string `yaml:"words"`
	ExpectRes int      `yaml:"expect_res"`
}

func TestSimilarPairs(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := similarPairs(caseData.Words)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
