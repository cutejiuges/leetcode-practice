package LC3307

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/7/4 下午11:33
 * @FilePath: kth_character_test.go
 * @Description:
 */

type TestData map[string]struct {
	K          int64  `yaml:"k"`
	Operations []int  `yaml:"operations"`
	ExpectRes  string `yaml:"expect_res"`
}

func TestKthCharacter(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := kthCharacter(caseData.K, caseData.Operations); !reflect.DeepEqual(res, caseData.ExpectRes[0]) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
