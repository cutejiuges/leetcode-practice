package LC2438

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/11 下午10:52
 * @FilePath: product_queries_test.go
 * @Description:
 */

type TestData map[string]struct {
	N         int     `yaml:"n"`
	Queries   [][]int `yaml:"queries"`
	ExpectRes []int   `yaml:"expect_res"`
}

func TestProductQueries(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := productQueries(caseData.N, caseData.Queries); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %vm ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
