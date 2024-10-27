package LC684

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/27 上午2:05
 * @FilePath: find_redundant_connection_test
 * @Description:
 */

type TestData map[string]struct {
	Edges     [][]int `yaml:"edges"`
	ExpectRes []int   `yaml:"expect_res"`
}

func TestFindRedundantConnection(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := findRedundantConnection(caseData.Edges)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
