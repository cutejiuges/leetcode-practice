package LC2545

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/21 下午8:40
 * @FilePath: sort_the_students_test
 * @Description:
 */

type TestData map[string]struct {
	Score     [][]int `yaml:"score"`
	K         int     `yaml:"k"`
	ExpectRes [][]int `yaml:"expect_res"`
}

func TestSortTheStudents(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := sortTheStudents(caseData.Score, caseData.K)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
