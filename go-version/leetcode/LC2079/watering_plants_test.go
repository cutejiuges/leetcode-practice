package LC2079

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/5/8 上午12:15
 * @FilePath: watering_plants_test
 * @Description:
 */

type TestData map[string]struct {
	Plants    []int `yaml:"Plants"`
	Capacity  int   `yaml:"Capacity"`
	ExpectRes int   `yaml:"ExpectRes"`
}

func TestWateringPlants(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := wateringPlants(caseData.Plants, caseData.Capacity)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
