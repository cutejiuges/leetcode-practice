package LC2125

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/10/27 下午11:28
 * @FilePath: number_of_beams_test.go
 * @Description:
 */

type TestData map[string]struct {
	Bank      []string `yaml:"bank"`
	ExpectRes int      `yaml:"expect_res"`
}

func TestNumberOfBeams(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := numberOfBeams(caseData.Bank); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
