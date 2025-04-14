package LC1534

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/14 上午9:08
 * @FilePath: count_good_triplets_test.go
 * @Description:
 */

type TestData map[string]struct {
	Arr       []int `yaml:"arr"`
	A         int   `yaml:"a"`
	B         int   `yaml:"b"`
	C         int   `yaml:"c"`
	ExpectRes int   `yaml:"expect_res"`
}

func TestCountGoodTriplets(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := countGoodTriplets(caseData.Arr, caseData.A, caseData.B, caseData.C)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
