package LC165

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/9/23 下午11:38
 * @FilePath: compare_version_test.go
 * @Description:
 */

type TestData map[string]struct {
	Version1  string `yaml:"version1"`
	Version2  string `yaml:"version2"`
	ExpectRes int    `yaml:"expect_res"`
}

func TestCompareVersion(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := compareVersion(caseData.Version1, caseData.Version2); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
