package LC2264

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/1/8 下午9:01
 * @FilePath: largest_good_integer_test
 * @Description:
 */

type TestData map[string]struct {
	Num       string `yaml:"num"`
	ExpectRes string `yaml:"expect_res"`
}

func TestLargestGoodInteger(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := largestGoodInteger(caseData.Num)
		if reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
