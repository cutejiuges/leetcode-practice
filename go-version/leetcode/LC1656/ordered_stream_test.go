package LC1656

import (
	"cutejiuge/leetcode-practice/util"
	"reflect"
	"testing"
)

/**
 * @Author: cutejiuge cutejiuge@163.com
 * @Date: 2025/2/24 下午11:13
 * @FilePath: ordered_stream_test
 * @Description:
 */

type TestData map[string]struct {
	Nums      []int      `yaml:"nums"`
	Strings   []string   `yaml:"strings"`
	ExpectRes [][]string `yaml:"expect_res"`
}

func TestConstructor(t *testing.T) {
	var testData TestData
	util.DataProvide(&testData)

	for caseName, caseData := range testData {
		var res [][]string
		stream := Constructor(caseData.Nums[0])
		for i, str := range caseData.Strings {
			res = append(res, stream.Insert(caseData.Nums[i+1], str))
		}
		if !reflect.DeepEqual(res, caseData.ExpectRes) {
			t.Errorf("%s, res = %v, ExpectRes = %v", caseName, res, caseData.ExpectRes)
		}
	}
}
