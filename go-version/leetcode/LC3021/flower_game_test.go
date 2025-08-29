package LC3021

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/8/29 下午11:27
 * @FilePath: flower_game_test.go
 * @Description:
 */

type TestData map[string]struct {
	N         int   `yaml:"n"`
	M         int   `yaml:"m"`
	ExpectRes int64 `yaml:"expect_res"`
}

func TestFlowerGame(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := flowerGame(caseData.N, caseData.M); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
