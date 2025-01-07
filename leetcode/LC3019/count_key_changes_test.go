package LC3019

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/1/7 下午11:30
 * @FilePath: count_key_changes_test
 * @Description:
 */

type TestData map[string]struct {
	S         string `yaml:"s"`
	ExpectRes int    `yaml:"expect_res"`
}

func TestCountKeyChanges(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := countKeyChanges(caseData.S)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
