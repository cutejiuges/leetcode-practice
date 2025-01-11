package LC3083

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/26 下午11:11
 * @FilePath: is_substring_present_test
 * @Description:
 */

type TestData map[string]struct {
	S         string `yaml:"s"`
	ExpectRes bool   `yaml:"expect_res"`
}

func TestIsSubstringPresent(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := isSubstringPresent(caseData.S)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
