package LC386

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/6/8 上午8:43
 * @FilePath: lexical_order_test.go
 * @Description:
 */

type TestData map[string]struct {
	N         int   `yaml:"n"`
	ExpectRes []int `yaml:"expect_res"`
}

func TestLexicalOrder(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := lexicalOrder(caseData.N); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
