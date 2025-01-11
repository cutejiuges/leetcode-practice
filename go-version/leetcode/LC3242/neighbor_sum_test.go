package LC3242

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"strconv"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/9 下午5:02
 * @FilePath: neighbor_sum_test
 * @Description:
 */

type TestData map[string]struct {
	Commands  []string `yaml:"commands"`
	Grid      [][]int  `yaml:"grid"`
	Values    []int    `yaml:"values"`
	ExpectRes []string `yaml:"expect_res"`
}

func TestNeighborSum(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := make([]string, 0)
		var nbs NeighborSum
		for i, ss := range caseData.Commands {
			switch ss {
			case "neighborSum":
				nbs = Constructor(caseData.Grid)
				res = append(res, "null")
			case "adjacentSum":
				res = append(res, strconv.Itoa(nbs.AdjacentSum(caseData.Values[i-1])))
			case "diagonalSum":
				res = append(res, strconv.Itoa(nbs.DiagonalSum(caseData.Values[i-1])))
			}
		}
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
