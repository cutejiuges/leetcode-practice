package LC75

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/17 下午9:58
 * @FilePath: sort_colors_test.go
 * @Description:
 */

type TestData map[string]struct {
	Nums      []int `yaml:"nums"`
	ExpectRes []int `yaml:"expect_res"`
}

func TestSortColors(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if sortColors(caseData.Nums); !reflect.DeepEqual(caseData.Nums, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, caseData.Nums, caseData.ExpectRes)
			}
		})
	}
}
