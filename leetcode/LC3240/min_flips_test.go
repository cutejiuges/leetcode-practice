package LC3240

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/16 下午9:13
 * @FilePath: min_flips_test
 * @Description:
 */

type TestData map[string]struct {
	Grid      [][]int `yaml:"grid"`
	ExpectRes int     `yaml:"expect_res"`
}

func TestMinFlips(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := minFlips(caseData.Grid)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
