package LC2476

import (
	datastruct "cutejiuge/leetcode-practice/data_struct"
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/7/10 上午12:17
 * @FilePath: closest_node_test
 * @Description:
 */

type TestData map[string]struct {
	Root      []string `yaml:"Root"`
	Queries   []int    `yaml:"Queries"`
	ExpectRes [][]int  `yaml:"ExpectRes"`
}

func TestClosestNodes(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		root := datastruct.Construct(caseData.Root)
		res := closestNodes(root, caseData.Queries)
		if len(res) != len(caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}

		for i, v := range res {
			if v[0] != caseData.ExpectRes[i][0] || v[1] != caseData.ExpectRes[i][1] {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		}
	}
}
