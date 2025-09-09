package LC2327

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/9/9 下午11:43
 * @FilePath: people_aware_of_secret_test.go
 * @Description:
 */

type TestData map[string]struct {
	N         int `yaml:"n"`
	Delay     int `yaml:"delay"`
	Forget    int `yaml:"forget"`
	ExpectRes int `yaml:"expect_res"`
}

func TestPeopleAwareOfSecret(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := peopleAwareOfSecret(caseData.N, caseData.Delay, caseData.Forget); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectREs = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
