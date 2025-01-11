package LC2391

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/11 上午12:29
 * @FilePath: garbage_collection_test
 * @Description:
 */

type TestData map[string]struct {
	Garbage   []string `yaml:"Garbage"`
	Travel    []int    `yaml:"Travel"`
	ExpectRes int      `yaml:"ExpectRes"`
}

func TestGarbageCollection(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := garbageCollection(caseData.Garbage, caseData.Travel)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
