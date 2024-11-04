package LC633

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/4 上午9:16
 * @FilePath: judge_square_sum_test
 * @Description:
 */

type TestData map[string]struct {
	C         int  `yaml:"c"`
	ExpectRes bool `yaml:"expect_res"`
}

func TestJudgeSquareSum(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := judgeSquareSum(caseData.C)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
