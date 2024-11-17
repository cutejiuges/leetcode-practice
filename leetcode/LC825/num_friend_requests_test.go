package LC825

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/11/17 下午2:08
 * @FilePath: num_friend_requests_test
 * @Description:
 */

type TestData map[string]struct {
	Ages      []int `yaml:"ages"`
	ExpectRes int   `yaml:"expect_res"`
}

func TestNumFriendRequests(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := numFriendRequests(caseData.Ages)
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
