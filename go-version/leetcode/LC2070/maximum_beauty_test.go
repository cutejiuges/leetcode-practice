package LC2070

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/9 下午4:08
 * @FilePath: maximum_beauty_test
 * @Description:
 */

type TestData map[string]struct {
	Items     [][]int `yaml:"items"`
	Queries   []int   `yaml:"queries"`
	ExpectRes []int   `yaml:"expect_res"`
}

func TestMaximumBeauty(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := maximumBeauty(caseData.Items, caseData.Queries)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
