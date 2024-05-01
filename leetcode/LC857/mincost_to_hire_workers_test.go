package LC857

import (
	"cutejiuge/leetcode-practice/util"
	"math"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/2 上午1:01
 * @FilePath: mincost_to_hire_workers_test
 * @Description:
 */

type TestData map[string]struct {
	Quality   []int   `yaml:"Quality"`
	Wage      []int   `yaml:"Wage"`
	K         int     `yaml:"K"`
	ExpectRes float64 `yaml:"ExpectRes"`
}

func TestMincostToHireWorkers(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)
	for caseName, caseData := range testData {
		res := mincostToHireWorkers(caseData.Quality, caseData.Wage, caseData.K)
		if math.Abs(res-caseData.ExpectRes) > 1e-5 {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
