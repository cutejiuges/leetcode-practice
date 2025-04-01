package LC2140

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/1 下午10:51
 * @FilePath: most_points_test.go
 * @Description:
 */

type TestData map[string]struct {
	Questions [][]int `yaml:"questions"`
	ExpectRes int64   `yaml:"expect_res"`
}

func TestMostPoints(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := mostPoints(caseData.Questions)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
