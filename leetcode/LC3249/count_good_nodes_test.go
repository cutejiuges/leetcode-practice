package LC3249

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/14 下午11:43
 * @FilePath: count_good_nodes_test
 * @Description:
 */

type TestData map[string]struct {
	Edges     [][]int `yaml:"edges"`
	ExpectRes int     `yaml:"expect_res"`
}

func TestCountGoodNodes(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := countGoodNodes(caseData.Edges)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
