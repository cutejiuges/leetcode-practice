package LC1338

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/15 下午8:25
 * @FilePath: min_set_size_test
 * @Description:
 */

type TestData map[string]struct {
	Arr       []int `yaml:"arr"`
	ExpectRes int   `yaml:"expect_res"`
}

func TestMinSetSize(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := minSetSize(caseData.Arr)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpecvtRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
