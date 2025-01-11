package LC1387

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/22 下午9:21
 * @FilePath: get_k_th_test
 * @Description:
 */

type TestData map[string]struct {
	Lo        int `yaml:"lo"`
	Hi        int `yaml:"hi"`
	K         int `yaml:"k"`
	ExpectRes int `yaml:"expect_res"`
}

func TestGetKth(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := getKth(caseData.Lo, caseData.Hi, caseData.K)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
