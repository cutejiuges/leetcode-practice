package LC551

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/8/18 下午8:02
 * @FilePath: check_record_test
 * @Description:
 */

type TestData map[string]struct {
	S         string `yaml:"S"`
	ExpectRes bool   `yaml:"ExpectRes"`
}

func TestCheckRecord(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := checkRecord(caseData.S)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
