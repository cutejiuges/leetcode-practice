package LC2234

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/8 下午12:47
 * @FilePath: maximum_beauty_test
 * @Description:
 */

type TestData map[string]struct {
	Flowers    []int `yaml:"flowers"`
	NewFlowers int64 `yaml:"new_flowers"`
	Full       int   `yaml:"full"`
	Target     int   `yaml:"target"`
	Partial    int   `yaml:"partial"`
	ExpectRes  int64 `yaml:"expect_res"`
}

func TestMaximumBeauty(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := maximumBeauty(caseData.Flowers, caseData.NewFlowers, caseData.Target, caseData.Full, caseData.Partial)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
