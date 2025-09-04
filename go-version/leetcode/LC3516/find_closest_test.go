package LC3516

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/9/4 下午11:42
 * @FilePath: find_closest_test.go
 * @Description:
 */

type TestData map[string]struct {
	X         int `yaml:"x"`
	Y         int `yaml:"y"`
	Z         int `yaml:"z"`
	ExpectRes int `yaml:"expect_res"`
}

func TestFindClosest(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := findClosest(caseData.X, caseData.Y, caseData.Z); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
