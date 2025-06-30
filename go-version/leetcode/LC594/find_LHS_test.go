package LC594

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/6/30 下午10:57
 * @FilePath: find_LHS_test.go
 * @Description:
 */

type testData map[string]struct {
	Nums      []int `yaml:"nums"`
	ExpectRes int   `yaml:"expect_res"`
}

func TestFindLHS(t *testing.T) {
	var testData testData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := findLHS(caseData.Nums); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
