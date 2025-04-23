package LC1399

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/23 下午10:07
 * @FilePath: count_largest_group_test.go
 * @Description:
 */

type TestData map[string]struct {
	N         int `yaml:"n"`
	ExpectRes int `yaml:"expect_res"`
}

func TestCountLargestGroup(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := countLargestGroup(caseData.N); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
