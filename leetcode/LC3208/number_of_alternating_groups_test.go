package LC3208

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/27 下午8:38
 * @FilePath: number_of_alternating_groups_test
 * @Description:
 */

type TestData map[string]struct {
	colors    []int `yaml:"colors"`
	k         int   `yaml:"k"`
	expectRes int   `yaml:"expect_res"`
}

func TestNumberOfAlternatingGroups(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := numberOfAlternatingGroups(caseData.colors, caseData.k)
		if !reflect.DeepEqual(res, caseData.expectRes) {
			t.Errorf("%s, res = %v, ExpectZRes = %v", caseName, res, caseData.expectRes)
		}
	}
}
