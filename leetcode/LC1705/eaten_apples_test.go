package LC1705

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/24 下午11:14
 * @FilePath: eaten_apples_test
 * @Description:
 */

type TestData map[string]struct {
	Apples    []int `yaml:"apples"`
	Days      []int `yaml:"days"`
	ExpectRes int   `yaml:"expect_res"`
}

func TestEatenApples(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := eatenApples(caseData.Apples, caseData.Days)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
