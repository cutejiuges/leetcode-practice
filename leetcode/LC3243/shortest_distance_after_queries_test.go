package LC3243

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/19 下午11:07
 * @FilePath: shortest_distance_after_queries_test
 * @Description:
 */

type TestData map[string]struct {
	N         int     `yaml:"n"`
	Queries   [][]int `yaml:"queries"`
	ExpectRes []int   `yaml:"expect_res"`
}

func TestShortestDistanceAfterQueries(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := shortestDistanceAfterQueries(caseData.N, caseData.Queries)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
