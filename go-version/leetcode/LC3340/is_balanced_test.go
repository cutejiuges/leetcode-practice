package LC3340

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/14 下午11:20
 * @FilePath: is_balanced_test
 * @Description:
 */

type TestData map[string]struct {
	Num       string `yaml:"num"`
	ExpectRes bool   `yaml:"expect_res"`
}

func TestIsBalance(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := isBalanced(caseData.Num)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
