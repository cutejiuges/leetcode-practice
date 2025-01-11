package LC1870

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/10/2 下午1:58
 * @FilePath: min_speed_on_time_test
 * @Description:
 */

type TestData map[string]struct {
	Dist      []int   `yaml:"dist"`
	Hour      float64 `yaml:"hour"`
	ExpectRes int     `yaml:"expect_res"`
}

func TestMinSpeedOnTime(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := minSpeedOnTime(caseData.Dist, caseData.Hour)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
