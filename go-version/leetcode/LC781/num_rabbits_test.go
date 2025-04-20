package LC781

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/4/20 上午8:11
 * @FilePath: num_rabbits_test.go
 * @Description:
 */

type TestData map[string]struct {
	Answer    []int `yaml:"answer"`
	ExpectRes int   `yaml:"expect_res"`
}

func TestNumRabbits(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := numRabbits(caseData.Answer)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
