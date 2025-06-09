package LC440

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/6/9 下午11:44
 * @FilePath: find_kth_number_test.go
 * @Description:
 */

type TestData map[string]struct {
	N         int `yaml:"n"`
	K         int `yaml:"k"`
	ExpectRes int `yaml:"expect_res"`
}

func TestFindKthNumber(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := findKthNumber(caseData.N, caseData.K); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
