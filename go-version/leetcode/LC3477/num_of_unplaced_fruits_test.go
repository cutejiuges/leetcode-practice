package LC3477

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/5 下午11:12
 * @FilePath: num_of_unplaced_fruits_test.go
 * @Description:
 */

type TestData map[string]struct {
	Fruits    []int `yaml:"fruits"`
	Baskets   []int `yaml:"baskets"`
	ExpectRes int   `yaml:"expect_res"`
}

func TestNumOfUnplacedFruits(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := numOfUnplacedFruits(caseData.Fruits, caseData.Baskets); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
