package LC2073

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/9/29 下午11:28
 * @FilePath: time_required_to_buy_test
 * @Description:
 */

type TestData map[string]struct {
	Tickets   []int `yaml:"Tickets"`
	K         int   `yaml:"K"`
	ExpectRes int   `yaml:"ExpectRes"`
}

func TestTimeRequiredToBuy(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := timeRequiredToBuy(caseData.Tickets, caseData.K)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
