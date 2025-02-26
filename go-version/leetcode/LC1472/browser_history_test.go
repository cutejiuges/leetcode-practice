package LC1472

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/2/26 下午11:32
 * @FilePath: browser_history_test
 * @Description:
 */

type TestData map[string]struct {
	Commands  []string        `yaml:"commands"`
	Input     [][]interface{} `yaml:"input"`
	ExpectRes []interface{}   `yaml:"expect_res"`
}

func TestBrowserHistory(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		res := []interface{}{nil}
		bh := Constructor(caseData.Input[0][0].(string))
		for i := 1; i < len(caseData.Commands); i++ {
			switch caseData.Commands[i] {
			case "visit":
				bh.Visit(caseData.Input[i][0].(string))
				res = append(res, nil)
			case "back":
				res = append(res, bh.Back(caseData.Input[i][0].(int)))
			case "forward":
				res = append(res, bh.Forward(caseData.Input[i][0].(int)))
			}
		}
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
