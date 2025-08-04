package LC904

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/4 下午11:26
 * @FilePath: total_fruit_test.go
 * @Description:
 */

type TestData map[string]struct {
	Fruits    []int `yaml:"fruits"`
	ExpectRes int   `yaml:"expect_res"`
}

func TestTotalFruit(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := totalFruit(caseData.Fruits); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
