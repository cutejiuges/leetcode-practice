package LC790

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/5/5 下午9:07
 * @FilePath: num_tilings_test.go
 * @Description:
 */

type TestData map[string]struct {
	N         int `yaml:"n"`
	ExpectRes int `yaml:"expect_res"`
}

func TestNumTilings(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := numTilings(caseData.N); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
