package LC2614

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/3/18 下午11:08
 * @FilePath: diagonal_prime_test
 * @Description:
 */

type TestData map[string]struct {
	Nums      [][]int `yaml:"nums"`
	ExpectRes int     `yaml:"expect_res"`
}

func TestDiagonalPrime(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			res := diagonalPrime(caseData.Nums)
			if !reflect.DeepEqual(res, caseData.ExpectRes) {
				t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
			}
		})
	}
}
