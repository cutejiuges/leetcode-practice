package LC1366

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/12/29 下午10:11
 * @FilePath: rank_teams_test
 * @Description:
 */

type TestData map[string]struct {
	Votes     []string `yaml:"votes"`
	ExpectRes string   `yaml:"expect_res"`
}

func TestRankTeams(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := rankTeams(caseData.Votes)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
