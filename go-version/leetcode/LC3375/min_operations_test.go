package LC3375

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/9 上午12:18
 * @FilePath: min_operations_test.go
 * @Description:
 */

type TestData map[string]struct {
	Nums      []int `yaml:"nums"`
	K         int   `yaml:"k"`
	ExpectRes int   `yaml:"expect_res"`
}

func TestMinOperations(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := minOperations(caseData.Nums, caseData.K)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
