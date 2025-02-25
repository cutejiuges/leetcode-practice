package LC2502

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/2/26 上午12:01
 * @FilePath: allocator_test
 * @Description:
 */

type TestData map[string]struct {
	Commands  []string `yaml:"commands"`
	Input     [][]int  `yaml:"input"`
	ExpectRes []int    `yaml:"expect_res"`
}

func TestAllocator(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := []int{}
		allocator := Constructor(caseData.Input[0][0])
		for i := 1; i < len(caseData.Commands); i++ {
			switch caseData.Commands[i] {
			case "allocate":
				res = append(res, allocator.Allocate(caseData.Input[i][0], caseData.Input[i][1]))
			case "freeMemory":
				res = append(res, allocator.FreeMemory(caseData.Input[i][0]))
			}
		}
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
