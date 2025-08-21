package LC1504

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/21 下午11:00
 * @FilePath: num_submat_test.go
 * @Description:
 */

type TestData map[string]struct {
	Mat       [][]int `yaml:"mat"`
	ExpectRes int     `yaml:"expect_res"`
}

func TestNumSubmat(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := numSubmat(caseData.Mat); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectREs = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
