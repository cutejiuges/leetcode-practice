package lc63

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

type TestData map[string]struct {
	ObstacleGrid [][]int `yaml:"obstacle_grid"`
	ExpectRes    int     `yaml:"expect_res"`
}

func TestUniquePathsWithObstacles(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := uniquePathsWithObstacles(caseData.ObstacleGrid)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
