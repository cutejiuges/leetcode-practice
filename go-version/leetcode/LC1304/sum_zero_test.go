package LC1304

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/9/7 下午10:25
 * @FilePath: sum_zero_test.go
 * @Description:
 */

type TestData map[string]struct {
	N int `yaml:"n"`
}

func TestSumZero(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		t.Run(caseName, func(t *testing.T) {
			ans := sumZero(caseData.N)
			var sum int
			for _, v := range ans {
				sum += v
			}
			if sum != 0 {
				t.Errorf("%s, ans = %v, sum = %v", caseName, ans, sum)
			}
		})
	}
}
