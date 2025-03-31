package LC2278

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/31 下午11:27
 * @FilePath: percentage_letter_test.go
 * @Description:
 */

type TestData map[string]struct {
	S         string `yaml:"s"`
	Letter    string `yaml:"letter"`
	ExpectRes int    `yaml:"expect_res"`
}

func Test_percentageLetter(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := percentageLetter(caseData.S, caseData.Letter[0])
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
