package LC1491

import (
	"cutejiuge/leetcode-practice/util"
	"math"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/3 下午1:17
 * @FilePath: average_test
 * @Description:
 */

type TestData map[string]struct {
	Salary    []int   `yaml:"Salary"`
	ExpectRes float64 `yaml:"ExpectRes"`
}

func TestAverage(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := average(caseData.Salary)
		if math.Abs(res-caseData.ExpectRes) > 1e-5 {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
