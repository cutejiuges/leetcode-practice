package LC2806

import (
	"cutejiuge/leetcode-practice/util"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2024/6/25 下午11:01
 * @FilePath: account_balance_after_purchase_test
 * @Description:
 */

type TestData map[string]struct {
	PurchaseAmount int `yaml:"PurchaseAmount"`
	ExpectRes      int `yaml:"ExpectRes"`
}

func TestAccountBalanceAfterPurchase(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := accountBalanceAfterPurchase(caseData.PurchaseAmount)
		if res != caseData.ExpectRes {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
