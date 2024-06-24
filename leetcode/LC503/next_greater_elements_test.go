package LC503

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/24 下午11:46
 * @FilePath: next_greater_elements_test
 * @Description:
 */

type TestData map[string]struct {
	Nums      []int `yaml:"Nums"`
	ExpectRes []int `yaml:"ExpectRes"`
}

func TestNextGreaterElements(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := nextGreaterElements(caseData.Nums)
		if len(res) != len(caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}

		for i := range res {
			if res[i] != caseData.ExpectRes[i] {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		}
	}
}
