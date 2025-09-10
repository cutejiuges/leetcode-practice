package LC1733

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/9/10 下午11:40
 * @FilePath: minimum_teachings_test.go
 * @Description:
 */

type TestData map[string]struct {
	N           int     `yaml:"n"`
	Languages   [][]int `yaml:"languages"`
	FriendShips [][]int `yaml:"friendships"`
	ExpectRes   int     `yaml:"expect_res"`
}

func TestMinimumTeachings(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := minimumTeachings(caseData.N, caseData.Languages, caseData.FriendShips); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
