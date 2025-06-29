package LC1498

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/6/29 下午10:59
 * @FilePath: num_subseq_test
 * @Description:
 */

type TestData map[string]struct {
	Nums      []int `yaml:"nums"`
	Target    int   `yaml:"target"`
	ExpectRes int   `yaml:"expect_res"`
}

func TestNumSubseq(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			if res := numSubseq(caseData.Nums, caseData.Target); !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
