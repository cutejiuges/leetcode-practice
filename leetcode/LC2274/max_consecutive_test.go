package LC2274

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/1/6 上午8:20
 * @FilePath: max_consecutive_test
 * @Description:
 */

type TestData map[string]struct {
	Bottom    int   `yaml:"bottom"`
	Top       int   `yaml:"top"`
	Special   []int `yaml:"special"`
	ExpectRes int   `yaml:"expect_res"`
}

func TestMaxConsecutive(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := maxConsecutive(caseData.Bottom, caseData.Top, caseData.Special)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
