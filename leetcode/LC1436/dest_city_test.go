package LC1436

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/8 上午8:54
 * @FilePath: dest_city_test
 * @Description:
 */

type TestData map[string]struct {
	Paths     [][]string `yaml:"paths"`
	ExpectRes string     `yaml:"expect_res"`
}

func TestDestCity(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := destCity(caseData.Paths)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
